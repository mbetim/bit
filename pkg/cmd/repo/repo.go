package repo

import (
	"github.com/mbetim/bit/pkg/cmd"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Work with Bitbucket repositories",
	Long:  "Work with Bitbucket repositories",
}

func init() {
	cmd.RootCmd.AddCommand(repoCmd)

	repoCmd.PersistentFlags().StringP("workspace", "w", "", "Get the repositories from this workspace")
	repoCmd.PersistentFlags().StringP("repo", "r", "", "Get the pull requests from this repository")
}
