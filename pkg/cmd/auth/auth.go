package auth

import (
	"github.com/mbetim/bit/pkg/cmd"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate bit with Bitbucket",
	Long:  `Authenticate bit with Bitbucket`,
}

func init() {
	cmd.RootCmd.AddCommand(authCmd)
}
