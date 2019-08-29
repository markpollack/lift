package commands

import (
	"github.com/cloudlift/lift/pkg/initializr"
	"github.com/spf13/cobra"
)

// InitializrNewOptions contains passed to the 'initializr new' command
type InitializrNewOptions struct {
	Dependencies string
	GroupID      string
	ArtifactID   string
	Path         string
}

var initializrNewOptions InitializrNewOptions

// InitializrNewCommand Create the command that will create a new project using initializr
func InitializrNewCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "new",
		Short: "CreateNewProject Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		RunE:  doNewCommand,
	}
	listCmd.Flags().StringVar(&initializrNewOptions.Dependencies, "dependencies", "web", "project dependencies to use")
	listCmd.Flags().StringVar(&initializrNewOptions.GroupID, "groupId", "io.example", "artifact group ID to use")
	listCmd.Flags().StringVar(&initializrNewOptions.ArtifactID, "artifactId", "webdemo", "artifact id to create")
	listCmd.Flags().StringVar(&initializrNewOptions.Path, "path", "", "directory to unzip project, default = working directory")
	listCmd.Args = cobra.NoArgs
	return listCmd
}

func doNewCommand(cmd *cobra.Command, args []string) error {
	request := initializr.InitializrRequest{
		Dependencies: initializrNewOptions.Dependencies,
		GroupID:      initializrNewOptions.GroupID,
		ArtifactID:   initializrNewOptions.ArtifactID,
		Path:         initializrNewOptions.Path,
	}
	return initializr.CreateNewProject(request)
}
