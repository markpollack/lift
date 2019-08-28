package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func NewPlatformCommand() *cobra.Command {
	platformCmd := &cobra.Command{
		Use:   "platform",
		Short: "Platform commands",
		Long:  `Commands related to platform operations`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	platformCmd.AddCommand(NewPlatformListCommand())
	return platformCmd
}
