package pr

import (
	"fmt"
	"strconv"

	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/mbetim/bit/pkg/utils"
	"github.com/spf13/cobra"
)

var prViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the title, body, and other information about a pull request",
	Long:  "Display the title, body, and other information about a pull request",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid Pull Request number")
			return
		}

		repo, workspace, err := bitbucket.GetRepoAndWorkspaceNameFromCurrentDir()
		if err != nil {
			fmt.Println(err)
			return
		}

		pr, err := bitbucket.GetPullRequestById(workspace, repo, prId)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Opening " + pr.Links.Html.Href + " in your browser")
		utils.OpenBrowser(pr.Links.Html.Href)
	},
}

func init() {
	prViewCmd.Flags().BoolP("browser", "b", false, "Open the PR in the browser")
	prViewCmd.MarkFlagRequired("browser")

	prCmd.AddCommand(prViewCmd)
}
