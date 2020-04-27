package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_Root(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOutput(b)
	rootCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "" {
		t.Fatalf("expected \"%s\" got \"%s\"", "", string(out))
	}
}
