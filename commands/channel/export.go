package commands_channel

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandChannelExport() *cobra.Command {
	channel_export := &cobra.Command{
		Use:   "export",
		Short: "Export channel private key.",
		Run:   HandleCommandChannelExport,
	}

	channel_export.Flags().String("channel_id", "", "(str) claim id of channel to export")
	channel_export.Flags().String("channel_name", "", "(str) name of channel to export")
	channel_export.Flags().String("account_id", "", "(str) one or more account ids for accounts to look in for channels, defaults to all accounts.")
	channel_export.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return channel_export
}

func HandleCommandChannelExport(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["channel_id"]
		if exists {
			cmd.Help()
			return
		}
		params["channel_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "channel_export", params)
}
