package usecases

import "file-analizer/pkg/logger"

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
	return nil
}
