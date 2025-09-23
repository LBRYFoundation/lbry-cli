package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletStatus() *cobra.Command {
	wallet_status := &cobra.Command{
		Use:   "status",
		Short: "Status of wallet including encryption/lock state.",
		Run:   HandleCommandWalletStatus,
	}

	wallet_status.Flags().String("wallet_id", "", "(str) status of specific wallet")

	return wallet_status
}

func HandleCommandWalletStatus(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["wallet_id"]
		if exists {
			cmd.Help()
			return
		}
		params["wallet_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_status", params)
}
