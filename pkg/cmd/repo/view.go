package repo

import (
	"fmt"

	"github.com/mbetim/bit/pkg/utils"
	"github.com/spf13/cobra"
)

var repoViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the information about a repository",
	Long:  "Display the information about a repository",
	Run: func(cmd *cobra.Command, _ []string) {
		repo, workspace := utils.GetRepoAndWorkspaceFromCmdOrConfig(cmd)

		if repo == "" {
			fmt.Println("A repository must be specified, either as argument or in the current directory")
			return
		}

		if workspace == "" {
			fmt.Println("A workspace must be specified, either as argument or in the config file")
			return
		}

		url := "https://bitbucket.org/" + workspace + "/" + repo
		fmt.Println("Opening " + url + " in your browser")
		utils.OpenBrowser(url)
	},
}

func init() {
	repoViewCmd.Flags().BoolP("browser", "b", false, "Open repository in the browser")
	repoViewCmd.MarkFlagRequired("browser")

	repoCmd.AddCommand(repoViewCmd)
}
