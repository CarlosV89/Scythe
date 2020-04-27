package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var accountID string
var token string

func init() {

	initCmd.Flags().StringVarP(&accountID, "accountId", "a", "", "Harvest API AccountId")
	initCmd.Flags().StringVarP(&token, "token", "t", "", "Harvest API Token")

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create or edit config file",
	Long:  `This command is used to edit or create a configuration file that holds Harvest credentials`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if accountID == "" {
			prompt := promptui.Prompt{
				Label: "Enter your Harvest Account ID:",
			}

			accountID, err = prompt.Run()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if token == "" {
			prompt := promptui.Prompt{
				Label: "Enter your Harvest Account token:",
			}

			token, err = prompt.Run()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		viper.Set("Harvest.AccountId", accountID)
		viper.Set("Harvest.Token", token)

		if err := viper.SafeWriteConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				err := viper.WriteConfigAs(fmt.Sprintf("%s/.scythe.yml", home))
				fmt.Println(err)

				if err != nil {
					fmt.Printf("Something went wrong saving file: %s", err)
				}
			}
		}
	},
}
