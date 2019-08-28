package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// PlatformCommand has no direct actions. It is a container for subcommands such as `platform list`
func PlatformCommand() *cobra.Command {
	platformCmd := &cobra.Command{
		Use:   "platform",
		Short: "Platform commands",
		Long:  `Commands related to platform operations`,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		},
	}
	platformCmd.AddCommand(PlatformListCommand())
	return platformCmd
}
