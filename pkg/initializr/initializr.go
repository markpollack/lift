package initializr

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

// InitializrRequest contains options supported to generate a project using https://github.com/spring-io/initializr
type InitializrRequest struct {
	Dependencies string
	GroupID      string
	ArtifactID   string
	Path         string
}

// InitializrResponse contains the response from a project generation request
type InitializrResponse struct {
	ContentType string
	Contents    []byte
	Filename    string
}

// CreateNewProject creates a new project using start.spring.io
func CreateNewProject(request InitializrRequest) error {
	resp, err := generate(request)
	if err != nil {
		return err
	}
	if request.Path == "" {
		workingDir, err := os.Getwd()
		if err != nil {
			return err
		}
		return unpack(resp.Contents, workingDir)
	}
	return unpack(resp.Contents, request.Path)

}

func generate(request InitializrRequest) (InitializrResponse, error) {

	u, err := url.Parse("https://start.spring.io/starter.zip")
	if err != nil {
		return InitializrResponse{}, err
	}

	q := u.Query()
	q.Set("dependencies", request.Dependencies)
	q.Set("groupId", request.GroupID)
	q.Set("artifactId", request.ArtifactID)
	q.Set("type", "maven-project")
	u.RawQuery = q.Encode()

	log.Debug("Initializr encoded URL: ", u.String())

	// default timeout is infinite...
	var httpClient = &http.Client{
		Timeout: time.Second * 30,
	}

	fmt.Println("Invoking Initializr service at https://start.spring.io")

	resp, err := httpClient.Get(u.String())
	if err != nil {
		return InitializrResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return InitializrResponse{}, err
	}

	var initialzrResponse InitializrResponse
	initialzrResponse.Contents = body
	initialzrResponse.Filename = resp.Header.Get("Content-Disposition")
	initialzrResponse.ContentType = resp.Header.Get("Content-Type")

	return initialzrResponse, nil

}

//func closeResponse(response *http.Response) {
//	err := response.Body.Close()
//	if err != nil {
//		log.Warn("Can't close http response.", err)
//	}
//}

func unpack(zipContents []byte, targetPath string) error {

	zipReader, err := zip.NewReader(bytes.NewReader(zipContents), int64(len(zipContents)))
	if err != nil {
		return err
	}

	// Ensure targetPath is created
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		if err != nil {
			return err
		}
		if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
			return err
		}
	}

	// Iterate through each file/dir found in
	for _, file := range zipReader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		extractedFilePath := filepath.Join(
			targetPath,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Debug("Creating directory:", extractedFilePath)
			if err := os.MkdirAll(extractedFilePath, file.Mode()); err != nil {
				return err
			}
		} else {
			// Extract regular file since not a directory
			log.Debug("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("Initializr zip file extracted to " + targetPath)
	return nil
}
