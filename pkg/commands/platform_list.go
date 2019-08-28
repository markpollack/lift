package commands

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Platform list",
	Long:  `Lists platforms available to deploy to`,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Alias", "Type", "Profile", "Description"})

		table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor})

		data := [][]string{
			[]string{"gke-sandbox-cschaefer", "ci-k8s", "kubernetes", "qa", "GKE"},
			[]string{"cf-sandbox-cschaefer", "ci-cf", "cloudfoundry", "qa", "CF"},
		}

		table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

		table.AppendBulk(data)
		table.Render()
	},
}

func init() {
	platformCmd.AddCommand(listCmd)
}
