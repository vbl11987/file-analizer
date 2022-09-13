package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type directory interface {
	LoadFile(filepath string) error
	GenerateDirOutput() ([]byte, error)
}

type DirContent struct {
	Files []File `json:"files"`
}

type File struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func NewDirContent() directory {
	return &DirContent{
		Files: make([]File, 0),
	}
}

// LoadFile add the file received in the filepath to the directory list
func (dir *DirContent) LoadFile(filepath string) error {
	fi, err := os.Stat(filepath)
	if err != nil {
		return fmt.Errorf("failed getting the file information: %v", err)
	}
	file := File{
		Name: fi.Name(),
		Size: int(fi.Size()),
	}
	dir.Files = append(dir.Files, file)

	return nil
}

// GenerateDirOutput generate the json with the files in the directory
func (dir *DirContent) GenerateDirOutput() ([]byte, error) {
	result, err := json.MarshalIndent(dir, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("failed to generate the json output: %v", err)
	}
	return result, nil
}
