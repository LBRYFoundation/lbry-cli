package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	for _, arg := range os.Args {
		if arg == "-V" || arg == "--version" {
			info, _ := debug.ReadBuildInfo()
			fmt.Println("LBRY CLI " + info.Main.Version)
			os.Exit(0)
		}
	}
	//TODO
	os.Exit(0)
}
