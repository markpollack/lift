package testutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TempDir(t *testing.T, subpath string) (string, func()) {
	t.Helper()
	// Create under standard temp directory a 'lift-tests' directory
	tempDir, err := ioutil.TempDir("", "lift-tests-")
	if err != nil {
		t.Error("Failed to create the tempDir: "+tempDir, err)
	}

	tempSubDir := filepath.Join(tempDir, subpath)
	err = os.MkdirAll(tempSubDir, os.ModePerm)

	if err != nil {
		t.Error("Failed to create the tempDir: "+tempSubDir, err)
	}
	t.Log("Created: " + tempSubDir)
	return tempSubDir, func() {
		// Remove parent directory as well
		err = os.RemoveAll(tempDir)
		if err != nil {
			t.Error("Failed to remove tempDir: "+tempDir, err)
		}
	}
}

func FileContents(t *testing.T, filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Could not read the file:" + filename)
	}
	return string(content)
}
