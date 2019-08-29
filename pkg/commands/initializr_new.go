package commands

import (
	"github.com/cloudlift/lift/pkg/initializr"
	"github.com/spf13/cobra"
)

// InitialzrNewCommand contains the fields needed to create a new project and store it on disk.
type InitialzrNewCommand struct {
	cli
	dependencies string
	groupID      string
	artifactID   string
	path         string
}

// NewInitialzrCommand initializes the cobra.Command field and returns a pointer to a new InitialzrNewCommand instance.
func NewInitialzrCommand() *InitialzrNewCommand {
	var newCmd InitialzrNewCommand
	newCmd.cmd = &cobra.Command{
		Use:   "new",
		Short: "CreateNewProject Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return newCmd.runInitNewCommand()
		},
	}
	newCmd.addFlags()
	return &newCmd
}

func (i *InitialzrNewCommand) addFlags() {
	i.cmd.Flags().StringVar(&i.dependencies, "dependencies", "", "project dependencies to use")
	i.cmd.Flags().StringVar(&i.groupID, "group-id", "", "artifact group ID to use")
	i.cmd.Flags().StringVar(&i.artifactID, "artifact-id", "", "artifact id to create")
	i.cmd.Flags().StringVar(&i.path, "path", "", "directory to unzip project, default = working directory")
	i.cmd.Args = cobra.NoArgs
}

func (i *InitialzrNewCommand) runInitNewCommand() error {
	request := initializr.InitializrRequest{
		Dependencies: i.dependencies,
		GroupID:      i.groupID,
		ArtifactID:   i.artifactID,
	}
	return initializr.CreateNewProject(request, i.path)
}
