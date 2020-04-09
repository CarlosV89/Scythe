package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daily bool
var hours int
var description string

func init() {
	sowCmd.Flags().IntVarP(&hours, "hours", "h", 0, "How many hours in time do you wish to log?")
	sowCmd.Flags().StringVarP(&description, "description", "d", "", "Small description")

	rootCmd.AddCommand(sowCmd)
}

func sow(cmd *cobra.Command, args []string) error {
	fmt.Printf("%v", viper.Get("Id"))
	return nil
}

var sowCmd = &cobra.Command{
	Use:   "sow",
	Short: "Register hours of work",
	Long:  `This command is used to register x ammounts of hours of work into a specific work day`,
	RunE:  sow,
}
