package pr

import (
	"github.com/mbetim/bit/pkg/cmd"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Work with Bitbucket pull requests",
	Long:  "Work with Bitbucket pull requests",
}

func init() {
	cmd.RootCmd.AddCommand(prCmd)

	prCmd.PersistentFlags().StringP("workspace", "w", "", "Get the repositories from this workspace")
	prCmd.PersistentFlags().StringP("repo", "r", "", "Get the pull requests from this repository")
}
