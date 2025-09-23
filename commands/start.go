package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func CreateCommandStart() *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Short: "Start LBRY Network interface.",
		Run:   HandleCommandStart,
	}

	return start
}

func HandleCommandStart(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	fmt.Println("The daemon cannot be started using the CLI.")
	os.Exit(1)
}
