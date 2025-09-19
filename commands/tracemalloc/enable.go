package commands_tracemalloc

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTraceMAllocEnable() *cobra.Command {
	tracemalloc_enable := &cobra.Command{
		Use:   "enable",
		Short: "Enable tracemalloc memory tracing",
		Run:   HandleCommandTraceMAllocEnable,
	}

	return tracemalloc_enable
}

func HandleCommandTraceMAllocEnable(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("tracemalloc_enable")
}
