package auth

import (
	"encoding/base64"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/mbetim/bit/pkg/config"
	"github.com/spf13/cobra"
)

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login in to bit",
	Long:  `Login in to bit`,
	Run: func(_ *cobra.Command, _ []string) {
		validate := func(label string) func(string) error {
			return func(input string) error {
				if len(input) == 0 {
					return fmt.Errorf("%s is required", label)
				}

				return nil
			}
		}

		prompt := promptui.Prompt{
			Label:    "Username",
			Validate: validate("Username"),
		}

		username, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v", err)
			return
		}

		prompt = promptui.Prompt{
			Label:    "App password",
			Mask:     '*',
			Validate: validate("App password"),
		}

		appPassword, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v", err)
			return
		}

		token := base64.StdEncoding.EncodeToString([]byte(username + ":" + appPassword))

		config.StoreToken(token)
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
}
