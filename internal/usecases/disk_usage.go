package usecases

import (
	"errors"
	"file-analizer/internal/models"
	"file-analizer/pkg/generator"
	"file-analizer/pkg/logger"
	"file-analizer/pkg/validator"
	"fmt"
	"path/filepath"
)

type loggerFactory interface {
	Create() logger.Log
}

type diskUsage struct {
	logger loggerFactory
}

func NewDiskUsage(log loggerFactory) *diskUsage {
	return &diskUsage{logger: log}
}

// Execute runs the process for the given mount point
func (manager *diskUsage) Execute(path string) error {
	log := manager.logger.Create()
	log.Info("Executing the process for the given mount point in %s", path)

	log.Debug("Validating the given mount point: %v", path)
	if !validator.IsPathValid(path) {
		log.Fatal("invalid mount point", errors.New("please check that the path is valid and if you have acccess to it"))
	}

	listFiles, err := generator.GetFileListName(path)
	if err != nil {
		log.Fatal("failed to geberate the list of files", err)
	}
	// Creating a directoy object to generate the list of files
	directory := models.NewDirContent()
	// Looping through the list of files and get the information
	for i := range listFiles {
		directory.LoadFile(filepath.Join(path, listFiles[i]))
	}

	//generating the output json
	result, err := directory.GenerateDirOutput()
	if err != nil {
		return fmt.Errorf("failed to generate the output json: %v", err)
	}

	fmt.Println(string(result))

	return nil
}
