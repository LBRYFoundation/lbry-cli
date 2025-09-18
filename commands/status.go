package commands

import "lbry/cli/rpc"
import "github.com/spf13/cobra"

func CreateCommandStatus() *cobra.Command {
	status := &cobra.Command{
		Use:   "status",
		Short: "Get daemon status",
		Run:   HandleCommandStatus,
	}

	return status
}

func HandleCommandStatus(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("status")
}
