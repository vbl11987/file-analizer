package cmd

import "github.com/spf13/cobra"

type managerProcessor interface {
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

func UsageCommand(p managerProcessor, log log) *cobra.Command {
	usageCommand := &cobra.Command{
		Use:   "usage",
		Short: "Use the passed arguments (mount point) to generate a list of the files with the disk usage.",
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Executing application for usage command...")
			log.Info("CLI arguments %s", args)
			if err := p.Execute(args[0]); err != nil {
				log.Fatal("executing the usage command", err)
			}
		},
	}
	return usageCommand
}
