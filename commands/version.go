package commands

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandVersion() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Get lbrynet API server version information",
		Run:   HandleCommandVersion,
	}

	return version
}

func HandleCommandVersion(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("version")
}
