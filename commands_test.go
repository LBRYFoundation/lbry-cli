package main

import (
	"testing"
)

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()

	if cmd == nil {
		t.Fatal("Returned command shouldn't be nil")
	}
}
