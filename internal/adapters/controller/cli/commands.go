package cli

import (
	"file-analizer/internal/adapters/controller/cli/cmd"
	"os"

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

	commands := rootCmd.Commands()
	var cmdFound bool
	for _, a := range commands {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			}
		}
	}

	if !cmdFound {
		args := append([]string{"usage"}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	return rootCmd.Execute()
}
