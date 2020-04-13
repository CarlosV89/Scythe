package mock

import (
	"github.com/spf13/cobra"
)

//CobraRunFunction is a type for mocking a function call
type CobraRunFunction func([]string)

//CobraCommand is a mock function for mocking a cobra command, it is intended to be passed into functions meant to go in aa Run call, we will see how this works out
func CobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "testing-sow",
		Short: "A test command for testing our application functions",
	}
}
