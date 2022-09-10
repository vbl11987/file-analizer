package main

import (
	"file-analizer/internal/app"
	"file-analizer/pkg/logger"
	"fmt"
)

func failOnError(err error, msg string) {
	logger.NewFactory("error").Create().Fatal(msg, err)
}

func main() {
	fmt.Println("Executing file-analizer")
	a, err := app.NewApp()
	if err != nil {
		failOnError(err, "could not create the application instance")
	}

	err = a.Run()
	if err != nil {
		failOnError(err, "failed executing the file-analizer")
	}
}
