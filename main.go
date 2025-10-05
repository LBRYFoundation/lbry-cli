package main

import (
	"fmt"
	"os"

	"lbry/cli/internal/commands"
)

func main() {
	os.Exit(run())
}

func run() int {
	rootCmd := commands.CreateCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
