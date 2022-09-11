package cli

import (
	"file-analizer/internal/adapters/controller/cli/cmd"

	"github.com/spf13/cobra"
)

type managerCommand interface {
	Execute(path string) error
}

type log interface {
	Info(messageFormat string, v ...interface{})
	Fatal(message string, err ...error)
}

func Start(mgr managerCommand, log log) error {
	rootCmd := &cobra.Command{
		Use:   "file-analizer-cli",
		Short: "Given a mount point the application will generate a list of files with the disk use in json bytes format.",
	}
	rootCmd.AddCommand(cmd.UsageCommand(mgr, log))

	return rootCmd.Execute()
}
