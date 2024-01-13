package utils

import (
	"github.com/mbetim/bit/pkg/bitbucket"
	"github.com/mbetim/bit/pkg/config"
	"github.com/spf13/cobra"
)

func GetRepoAndWorkspaceFromCmdOrConfig(cmd *cobra.Command) (string, string) {
	workspace, _ := cmd.Flags().GetString("workspace")
	repo, _ := cmd.Flags().GetString("repo")

	if repo == "" {
		repo, workspace, _ = bitbucket.GetRepoAndWorkspaceNameFromCurrentDir()
		return repo, workspace
	}

	if workspace == "" {
		savedConfig, _ := config.GetConfig()

		workspace = savedConfig.DefaultWorkspace
	}

	return repo, workspace
}
