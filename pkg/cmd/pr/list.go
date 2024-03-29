package pr

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/mbetim/bit/pkg/utils"
	"github.com/spf13/cobra"
)

var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pull requests in a Bitbucket repository",
	Long:  `List pull requests in a Bitbucket repository`,
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

		prs, err := bitbucket.GetPullRequestsFromRepo(workspace, repo)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		cCyan := "\033[36m"
		cBlue := "\033[34m"
		reset := "\033[0m"

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprint(w, cCyan+"ID\t"+reset+"TITLE\t"+cBlue+"BRANCH"+reset+"\n")
		for _, pr := range prs {
			fmt.Fprintf(w, cCyan+"%v\t"+reset+"%s\t"+cBlue+"%s"+reset+"\n", pr.Id, pr.Title, pr.Source.Branch.Name)
		}
		w.Flush()
	},
}

func init() {
	prCmd.AddCommand(prListCmd)
}
