package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletLock() *cobra.Command {
	wallet_lock := &cobra.Command{
		Use:   "lock",
		Short: "Lock an unlocked wallet",
		Run:   HandleCommandWalletLock,
	}

	wallet_lock.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return wallet_lock
}

func HandleCommandWalletLock(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("wallet_lock", params)
}
