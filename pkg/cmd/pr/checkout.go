package pr

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/spf13/cobra"
)

var prCheckoutCmd = &cobra.Command{
	Use:   "checkout <id>",
	Short: "Check out a pull request in Bitbucket",
	Long:  `Check out a pull request in Bitbucket`,
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

		commands := []string{
			"git fetch origin",
			"git checkout " + pr.Source.Branch.Name,
			"git pull",
		}

		commandOutput, err := exec.Command("sh", "-c", strings.Join(commands, " && ")).CombinedOutput()
		if err != nil {
			fmt.Println(pr, err)
			return
		}

		fmt.Println(string(commandOutput))
	},
}

func init() {
	// TODO: Remove pr and workspace flags from this command
	prCmd.AddCommand(prCheckoutCmd)
}
