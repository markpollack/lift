package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lift-go",
	Short: "Multi cloud deployment tool",
	Long:  `lift is a tool for enriching your application so it can be deployed to multiple cloud platforms with minimal effort.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
