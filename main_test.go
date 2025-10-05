package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	exitCode := run()

	if exitCode != 0 {
		t.Fatal("Exit code should be zero.")
	}
}
