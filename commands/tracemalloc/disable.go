package commands_tracemalloc

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTraceMAllocDisable() *cobra.Command {
	tracemalloc_disable := &cobra.Command{
		Use:   "disable",
		Short: "Disable tracemalloc memory tracing",
		Run:   HandleCommandTraceMAllocDisable,
	}

	return tracemalloc_disable
}

func HandleCommandTraceMAllocDisable(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "tracemalloc_disable")
}
