package commands

import "github.com/spf13/cobra"

type cli struct {
	cmd *cobra.Command
}

func (c *cli) Cmd() *cobra.Command {
	return c.cmd
}
