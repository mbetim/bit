package pr

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/mbetim/bit/pkg/bitbucket"
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

		var shellCmd *exec.Cmd
		url := pr.Links.Html.Href

		switch runtime.GOOS {
		case "windows":
			shellCmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		case "darwin":
			shellCmd = exec.Command("open", url)
		case "linux":
			shellCmd = exec.Command("xdg-open", url)
		default:
			return
		}

		fmt.Println("Opening " + pr.Links.Html.Href + " in your browser")
		shellCmd.Start()
	},
}

func init() {
	prViewCmd.Flags().BoolP("browser", "b", false, "Open the PR in the browser")
	prViewCmd.MarkFlagRequired("browser")

	prCmd.AddCommand(prViewCmd)
}
