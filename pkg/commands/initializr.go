package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// InitializrCommand creates a new top level "initializr" command
func InitializrCommand() *cobra.Command {
	platformCmd := &cobra.Command{
		Use:   "initializr",
		Short: "Initializr commands",
		Long:  `Commands related to Spring Initializr operations`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(1)
		},
	}
	platformCmd.AddCommand(InitializrNewCommand())
	return platformCmd
}
