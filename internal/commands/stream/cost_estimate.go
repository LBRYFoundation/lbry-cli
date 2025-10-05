package commands_stream

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStreamCostEstimate() *cobra.Command {
	stream_cost_estimate := &cobra.Command{
		Use:   "cost_estimate",
		Short: "Get estimated cost for a lbry stream",
		Run:   HandleCommandStreamCostEstimate,
	}

	stream_cost_estimate.Flags().String("uri", "", "(str) uri to use")

	return stream_cost_estimate
}

func HandleCommandStreamCostEstimate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "uri")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["uri"]
		if exists {
			cmd.Help()
			return
		}
		params["uri"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "stream_cost_estimate", params)
}
