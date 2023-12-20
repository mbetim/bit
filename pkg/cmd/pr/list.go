package pr

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/spf13/cobra"
)

// prListCmd represents the list command
var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) == 0 || args[0] == "" {
		// 	fmt.Println("A repository must be specified")
		// 	return
		// }

		workspace, _ := cmd.Flags().GetString("workspace")

		// TODO: Get this value from a flag
		repo := ""
		if len(args) > 0 {
			repo = args[0]
		}

		if repo == "" {
			repo, _ = bitbucket.GetRepoNameFromCurrentDir()

			if repo == "" {
				fmt.Println("A repository must be specified, either as argument or in the current directory")
				return
			}
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	prListCmd.Flags().StringP("workspace", "w", "synvia", "Get the repositories from this workspace")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
