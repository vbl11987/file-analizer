package validator

import (
	"io/ioutil"
	"os"
)

func IsPathValid(path string) bool {
	fileName := path + "/" + "testFile.txt"
	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		return true
	}

	// Attempt to create a new file
	var d []byte
	if err := ioutil.WriteFile(fileName, d, 0644); err == nil {
		os.Remove(fileName)
		return true
	}

	return false
}
