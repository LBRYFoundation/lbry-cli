package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()

	if cmd == nil {
		t.Fatal("Returned command shouldn't be nil")
	}
}

func TestHandleCommand(t *testing.T) {
	cmd := &cobra.Command{}

	HandleCommand(cmd, nil)
}
