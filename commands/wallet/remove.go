package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletRemove() *cobra.Command {
	wallet_remove := &cobra.Command{
		Use:   "remove",
		Short: "Remove an existing wallet.",
		Run:   HandleCommandWalletRemove,
	}

	wallet_remove.Flags().String("wallet_id", "", "(str) name of wallet to remove")

	return wallet_remove
}

func HandleCommandWalletRemove(cmd *cobra.Command, args []string) {
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

	rpc.ExecuteRPCCommand("wallet_remove", params)
}
