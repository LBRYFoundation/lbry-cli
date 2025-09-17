package main

import "fmt"
import "os"

func main() {
	rootCmd := CreateCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
