package commands

import (
	"github.com/cloudlift/lift/pkg/initializr"
	"github.com/spf13/cobra"
)

type InitialzrNewCommand struct {
	baseCommand
	dependencies string
	groupId      string
	artifactId   string
	path         string
}

func (i *InitialzrNewCommand) Init() {
	i.cmd = &cobra.Command{
		Use:   "new",
		Short: "New Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.runInitNewCommand(args)
		},
	}
	i.addFlags()
}

func (i *InitialzrNewCommand) addFlags() {
	i.cmd.Flags().StringVar(&i.dependencies, "dependencies", "web", "project dependencies to use")
	i.cmd.Flags().StringVar(&i.groupId, "groupId", "io.example", "artifact group ID to use")
	i.cmd.Flags().StringVar(&i.artifactId, "artifactId", "webdemo", "artifact id to create")
	i.cmd.Flags().StringVar(&i.path, "path", "", "directory to unzip project, default = working directory")
	i.cmd.Args = cobra.NoArgs
}

func (i *InitialzrNewCommand) runInitNewCommand(args []string) error {
	request := initializr.InitializrRequest{
		Dependencies: i.dependencies,
		GroupId:      i.groupId,
		ArtifactId:   i.artifactId,
		Path:         i.path,
	}
	return initializr.New(request)
}
