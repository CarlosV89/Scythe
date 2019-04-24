package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daily bool
var hours int
var getGitCommitLogs bool

func init() {

	initCmd.Flags().IntVarP(&daily, "hours", "h", "", "How many hours in time do you wish to log")
	initCmd.Flags().StringVarP(&hours, "hours", "h", "", "How many hours in time do you wish to log")
	initCmd.Flags().StringVarP(&token, "git-logs", "g", "", "Whether you want your commit logs in your daily hours report")

	rootCmd.AddCommand(sowCmd)
}

var sowCmd = &cobra.Command{
	Use:   "sow",
	Short: "Register hours of work",
	Long:  `This command is used to register x ammounts of hours of work into a specific work day`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v", viper.Get("Id"))

	},
}
