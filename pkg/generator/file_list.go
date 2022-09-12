package generator

import (
	"os"
)

func GetFileListName(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() {
			result = append(result, f.Name())
		}
	}
	return result, nil
}
