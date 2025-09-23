package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletUnlock() *cobra.Command {
	wallet_unlock := &cobra.Command{
		Use:   "unlock",
		Short: "Unlock an encrypted wallet",
		Run:   HandleCommandWalletUnlock,
	}

	wallet_unlock.Flags().String("password", "", "(str) password to use for unlocking")
	wallet_unlock.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return wallet_unlock
}

func HandleCommandWalletUnlock(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "password")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["password"]
		if exists {
			cmd.Help()
			return
		}
		params["password"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("wallet_unlock", params)
}
