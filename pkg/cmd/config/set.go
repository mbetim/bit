package config

import (
	"fmt"

	"github.com/mbetim/bit/pkg/config"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update configuration with a value for the given key",
	Long:  `Update configuration with a value for the given key`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("set requires 2 arguments")
			return
		}

		key := args[0]
		value := args[1]

		switch key {
		case "default_workspace":
			if err := config.AddConfig(key, value); err != nil {
				fmt.Println("Error adding config: ", err)
			}
		default:
			fmt.Printf("%v is a unknown key\n", key)
		}
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
