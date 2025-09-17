package commands

import commands "lbry/cli/commands/stream"
import "github.com/spf13/cobra"

func CreateCommandStream() *cobra.Command {
	stream := &cobra.Command{
		Use:   "stream",
		Short: "Create, update, abandon, list and inspect your stream claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	stream.AddCommand(commands.CreateCommandStreamAbandon())
	stream.AddCommand(commands.CreateCommandStreamCostEstimate())
	stream.AddCommand(commands.CreateCommandStreamCreate())
	stream.AddCommand(commands.CreateCommandStreamList())
	stream.AddCommand(commands.CreateCommandStreamRepost())
	stream.AddCommand(commands.CreateCommandStreamUpdate())

	return stream
}
