package commands_wallet

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletAdd() *cobra.Command {
	wallet_add := &cobra.Command{
		Use:   "add",
		Short: "Add existing wallet.",
		Run:   HandleCommandWalletAdd,
	}

	wallet_add.Flags().String("wallet_id", "", "(str) wallet file name")

	return wallet_add
}

func HandleCommandWalletAdd(cmd *cobra.Command, args []string) {
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_add", params)
}
