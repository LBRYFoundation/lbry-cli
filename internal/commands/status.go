package commands

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStatus() *cobra.Command {
	status := &cobra.Command{
		Use:   "status",
		Short: "Get daemon status",
		Run:   HandleCommandStatus,
	}

	return status
}

func HandleCommandStatus(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "status")
}
