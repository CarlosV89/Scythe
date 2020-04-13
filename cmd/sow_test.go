package cmd

import (
	"testing"

	"github.com/vanor89/scythe/internal/platform/testing/mock"
)

func Test_Sow(t *testing.T) {
	mockCmd := mock.CobraCommand()

	args := []string{"--arg", "test-arg"}
	if err := sow(mockCmd, args); err != nil {
		t.Fatal(err)
	}
}
