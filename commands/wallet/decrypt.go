package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletDecrypt() *cobra.Command {
	wallet_decrypt := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt an encrypted wallet, this will remove the wallet password. The wallet must be unlocked to decrypt it",
		Run:   HandleCommandWalletDecrypt,
	}

	wallet_decrypt.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return wallet_decrypt
}

func HandleCommandWalletDecrypt(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("wallet_decrypt", params)
}
