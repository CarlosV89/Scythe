package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_Root(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOutput(b)
	rootCmd.SetArgs([]string{"--in", "testisawesome"})
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "Error: unknown flag: --in" {
		t.Fatalf("expected \"%s\" got \"%s\"", "Error: unknown flag: --in", string(out))
	}
}
