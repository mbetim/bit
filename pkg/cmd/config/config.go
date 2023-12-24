package config

import (
	"github.com/mbetim/bit/pkg/cmd"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Display or change configuration settings for bit",
	Long:  `Display or change configuration settings for bit`,
}

func init() {
	cmd.RootCmd.AddCommand(configCmd)
}
