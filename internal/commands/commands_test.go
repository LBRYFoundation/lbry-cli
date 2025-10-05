package commands

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // run the function that writes to stdout

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()

	if cmd == nil {
		t.Fatal("Returned command shouldn't be nil")
	}
}

func TestHandleCommand(t *testing.T) {
	helpCalled := false

	cmd := &cobra.Command{}
	cmd.Flags().Bool("version", false, "")
	cmd.SetHelpFunc(func(*cobra.Command, []string) {
		helpCalled = true
	})

	HandleCommand(cmd, nil)
	if !helpCalled {
		t.Error("Help should have been called if --version is absent")
	}

	cmd.Flags().Set("version", "true")
	output := captureStdout(func() {
		HandleCommand(cmd, nil)
	})
	if !strings.HasPrefix(output, "LBRY CLI") {
		t.Error("Version should start with 'LBRY CLI'.")
	}
}
