package commands

import (
	"github.com/cloudlift/lift/pkg/initializr"
	"github.com/spf13/cobra"
)

var (
	dependencies string
	groupId      string
	artifactId   string
	path         string
)

func InitializrNewCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "new",
		Short: "New Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		RunE:  doNewCommand,
	}
	listCmd.Flags().StringVar(&dependencies, "dependencies", "web", "project dependencies to use")
	listCmd.Flags().StringVar(&groupId, "groupId", "io.example", "artifact group ID to use")
	listCmd.Flags().StringVar(&artifactId, "artifactId", "webdemo", "artifact id to create")
	listCmd.Flags().StringVar(&path, "path", "", "directory to unzip project, default = working directory")
	listCmd.Args = cobra.NoArgs
	return listCmd
}

func doNewCommand(cmd *cobra.Command, args []string) error {
	request := initializr.InitializrRequest{
		Dependencies: dependencies,
		GroupId:      groupId,
		ArtifactId:   artifactId,
		Path:         path,
	}
	return initializr.New(request)
}
