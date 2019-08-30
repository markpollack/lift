package commands

import (
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "lift",
		Short: "multi cloud code generation tool",
		Long:  `lift is a tool for enriching your application so it can be deployed to multiple cloud platforms with minimal effort.`,
	}
	rootCmd.AddCommand(PlatformCommand())
	rootCmd.AddCommand(InitializrCommand())
	return rootCmd
}
