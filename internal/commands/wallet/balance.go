package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletBalance() *cobra.Command {
	wallet_balance := &cobra.Command{
		Use:   "balance",
		Short: "Return the balance of a wallet",
		Run:   HandleCommandWalletBalance,
	}

	wallet_balance.Flags().String("wallet_id", "", "(str) balance for specific wallet")
	wallet_balance.Flags().Int("confirmations", -1, "(int) Only include transactions with this many confirmed blocks.")

	return wallet_balance
}

func HandleCommandWalletBalance(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "confirmations")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_balance", params)
}
