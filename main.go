package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	rootCmd := CreateCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
