package rpc

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/spf13/pflag"
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

func TestAddParameter(t *testing.T) {
	params := map[string]any{}

	flagSet := &pflag.FlagSet{}
	flagSet.String("key", "", "")
	flagSet.Set("key", "value")

	AddParameter(params, flagSet, flagSet.GetString, "key")

	if params["key"] != "value" {
		t.Error("Key should be in map with correct value.")
	}
}

func TestGetURL(t *testing.T) {
	// TODO
}

func TestExecuteRPCCommand(t *testing.T) {
	// TODO
}

func TestPrintRawJSON(t *testing.T) {
	output := ""

	output = captureStdout(func() {
		printRawJSON(true)
	})
	if output != "true\n" {
		t.Error("Function should print correct JSON.")
	}

	output = captureStdout(func() {
		printRawJSON(false)
	})
	if output != "false\n" {
		t.Error("Function should print correct JSON.")
	}

	output = captureStdout(func() {
		printRawJSON("string")
	})
	if output != "\"string\"\n" {
		t.Error("Function should print correct JSON.")
	}
}
