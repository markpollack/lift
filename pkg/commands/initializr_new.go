package commands

import (
	"os"

	"github.com/cloudlift/lift/pkg/initializr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type InitializrNewOptions struct {
	ArtifactId string
}

var flags InitializrNewOptions

func InitializrNewCommand() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "new",
		Short: "New Spring Boot application",
		Long:  `Creates a new Spring Boot application using Initializr`,
		Run: func(cmd *cobra.Command, args []string) {
			request := initializr.InitializrRequest{}
			artifactId, _ := cmd.Flags().GetString("artifactid")
			if artifactId != "" {
				request.ArtifactId = artifactId
			}
			resp, err := initializr.Generate(request)
			if err != nil {
				log.Fatal(err)
			}
			workingDir, err := os.Getwd()
			initializr.Unpack(resp.Contents, workingDir)
		},
	}
	listCmd.Flags().StringVar(&flags.ArtifactId, "artifactid", "webdemo", "artifact id to create")
	return listCmd
}
