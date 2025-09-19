package commands

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStop() *cobra.Command {
	stop := &cobra.Command{
		Use:   "stop",
		Short: "Stop lbrynet API server.",
		Run:   HandleCommandStop,
	}

	return stop
}

func HandleCommandStop(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("stop")
}
