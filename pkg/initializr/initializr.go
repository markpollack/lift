package initializr

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

type InitializrRequest struct {
	ArtifactId string
}

type InitializrResponse struct {
	ContentType string
	Contents    []byte
	Filename    string
}

func Generate(request InitializrRequest) (InitializrResponse, error) {

	baseUrl := "https://start.spring.io/starter.zip"

	u, err := url.Parse(baseUrl)
	q := u.Query()
	q.Set("dependencies", "web")
	q.Set("groupId", "io.example")
	q.Set("artifactId", request.ArtifactId)
	q.Set("type", "maven-project")
	u.RawQuery = q.Encode()

	log.Info("Encoded URL is ", u.String())

	// default timeout is infinite...
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(u.String())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var initialzrResponse InitializrResponse
	initialzrResponse.Contents = body
	initialzrResponse.Filename = resp.Header.Get("Content-Disposition")
	initialzrResponse.ContentType = resp.Header.Get("Content-Type")

	return initialzrResponse, nil

}

func Unpack(zipContents []byte, targetPath string) {

	zipReader, err := zip.NewReader(bytes.NewReader(zipContents), int64(len(zipContents)))
	if err != nil {
		log.Fatal(err)
	}

	// Ensure targetPath is created
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		os.MkdirAll(targetPath, os.ModePerm)
	}

	// Iterate through each file/dir found in
	for _, file := range zipReader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
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
			os.MkdirAll(extractedFilePath, file.Mode())
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
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
