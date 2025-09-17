package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamCostEstimate() *cobra.Command {
	stream_cost_estimate := &cobra.Command{
		Use:   "cost_estimate",
		Short: "Get estimated cost for a lbry stream",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_cost_estimate
}
