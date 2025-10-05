package commands

import (
	commands "lbry/cli/commands/stream"

	"github.com/spf13/cobra"
)

func CreateCommandStream() *cobra.Command {
	stream := &cobra.Command{
		Use:   "stream",
		Short: "Create, update, abandon, list and inspect your stream claims.",
	}

	stream.AddCommand(commands.CreateCommandStreamAbandon())
	stream.AddCommand(commands.CreateCommandStreamCostEstimate())
	stream.AddCommand(commands.CreateCommandStreamCreate())
	stream.AddCommand(commands.CreateCommandStreamList())
	stream.AddCommand(commands.CreateCommandStreamRepost())
	stream.AddCommand(commands.CreateCommandStreamUpdate())

	return stream
}
