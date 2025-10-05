package commands_channel

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandChannelImport() *cobra.Command {
	channel_import := &cobra.Command{
		Use:   "import",
		Short: "Import serialized channel private key (to allow signing new streams to the channel)",
		Run:   HandleCommandChannelImport,
	}

	channel_import.Flags().String("channel_data", "", "(str) serialized channel, as exported by channel export")
	channel_import.Flags().String("wallet_id", "", "(str) import into specific wallet")

	return channel_import
}

func HandleCommandChannelImport(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_data")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["channel_data"]
		if exists {
			cmd.Help()
			return
		}
		params["channel_data"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "channel_import", params)
}
