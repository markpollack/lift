package commands

import (
	"os"

	"github.com/cloudlift/lift/pkg/initializr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type InitializrNewOptions struct {
	Dependencies string
	GroupId      string
	ArtifactId   string
	Path         string
}

var flags InitializrNewOptions

func InitializrNewCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "new",
		Short: "New Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		Run: func(cmd *cobra.Command, args []string) {
			request := initializr.InitializrRequest{}
			request.Dependencies, _ = cmd.Flags().GetString("dependencies")
			request.GroupId, _ = cmd.Flags().GetString("groupId")
			request.ArtifactId, _ = cmd.Flags().GetString("artifactId")
			resp, err := initializr.Generate(request)
			if err != nil {
				log.Fatal(err)
			}
			if requestPath, _ := cmd.Flags().GetString("path"); requestPath == "" {
				workingDir, _ := os.Getwd()
				initializr.Unpack(resp.Contents, workingDir)
			} else {
				initializr.Unpack(resp.Contents, requestPath)
			}
		},
	}
	listCmd.Flags().StringVar(&flags.Dependencies, "dependencies", "web", "project dependencies to use")
	listCmd.Flags().StringVar(&flags.GroupId, "groupId", "io.example", "artifact group ID to use")
	listCmd.Flags().StringVar(&flags.ArtifactId, "artifactId", "webdemo", "artifact id to create")
	listCmd.Flags().StringVar(&flags.ArtifactId, "path", "", "directory to unzip project, default = working directory")

	listCmd.Args = cobra.NoArgs
	return listCmd
}
