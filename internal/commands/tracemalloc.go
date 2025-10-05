package commands

import (
	commands "lbry/cli/internal/commands/tracemalloc"

	"github.com/spf13/cobra"
)

func CreateCommandTraceMAlloc() *cobra.Command {
	tracemalloc := &cobra.Command{
		Use:   "tracemalloc",
		Short: "Controls and queries tracemalloc memory tracing tools for troubleshooting.",
	}

	tracemalloc.AddCommand(commands.CreateCommandTraceMAllocDisable())
	tracemalloc.AddCommand(commands.CreateCommandTraceMAllocEnable())
	tracemalloc.AddCommand(commands.CreateCommandTraceMAllocTop())

	return tracemalloc
}
