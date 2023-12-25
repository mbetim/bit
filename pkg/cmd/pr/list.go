package pr

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/mbetim/bit/pkg/config"
	"github.com/spf13/cobra"
)

var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pull requests in a Bitbucket repository",
	Long:  `List pull requests in a Bitbucket repository`,
	Run: func(cmd *cobra.Command, _ []string) {
		workspace, _ := cmd.Flags().GetString("workspace")
		repo, _ := cmd.Flags().GetString("repo")

		if repo != "" {
			if workspace != "" {
				savedConfig, _ := config.GetConfig()

				workspace = savedConfig.DefaultWorkspace
			}
		} else {
			repo, workspace, _ = bitbucket.GetRepoAndWorkspaceNameFromCurrentDir()
		}

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
			fmt.Printf("Error: %v", err)
			os.Exit(1)
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

	// prListCmd.Flags().StringP("workspace", "w", "synvia", "Get the repositories from this workspace")
	prListCmd.Flags().StringP("workspace", "w", "", "Get the repositories from this workspace")
	prListCmd.Flags().StringP("repo", "r", "", "Get the pull requests from this repository")
}
