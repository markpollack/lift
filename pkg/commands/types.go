package commands

import "github.com/spf13/cobra"

type Command interface {
	Init()
	Cmd() *cobra.Command
}

type baseCommand struct {
	cmd *cobra.Command
}

func (b *baseCommand) Init() {}

func (b *baseCommand) Cmd() *cobra.Command {
	return b.cmd
}
