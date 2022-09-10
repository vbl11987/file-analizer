package cli

import (
	"file-analizer/internal/adapters/controller/cli/cmd"

	"github.com/spf13/cobra"
)

type managerCommand interface {
	Execute(path string) error
}

type log interface {
	Debug(messageFormat string, v ...interface{})
	Info(messageFormat string, v ...interface{})
	Error(message string, err ...error)
	Fatal(message string, err ...error)
	InfoWithFields(message string, fields map[string]interface{}, err ...error)
	AddFieldToContext(message string, field interface{})
}

func Start(mgr managerCommand, log log) error {
	rootCmd := &cobra.Command{
		Use:   "file-analizer-cli",
		Short: "Given a mount point the application will generate a list of files with the disk use in json bytes format.",
	}
	rootCmd.AddCommand(cmd.UsageCommand(mgr, log))

	return rootCmd.Execute()
}
