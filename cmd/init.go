package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Get user information required for API Interaction",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("")
  },
}
