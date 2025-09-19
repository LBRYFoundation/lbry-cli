package commands_tracemalloc

import "github.com/spf13/cobra"

func CreateCommandTraceMAllocDisable() *cobra.Command {
	tracemalloc_disable := &cobra.Command{
		Use:   "disable",
		Short: "Disable tracemalloc memory tracing",
		Run:   HandleCommandTraceMAllocDisable,
	}

	return tracemalloc_disable
}

func HandleCommandTraceMAllocDisable(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("tracemalloc_disable")
}
