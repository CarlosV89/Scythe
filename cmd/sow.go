package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(sowCmd)
}

var sowCmd = &cobra.Command{
	Use:   "sow",
	Short: "Register hours of work",
	Long:  `This command is used to register x ammounts of hours of work into a specific work day`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		fmt.Printf("%v", viper.Get("Id"))
	},
}
