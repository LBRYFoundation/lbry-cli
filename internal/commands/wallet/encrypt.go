package commands_wallet

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletEncrypt() *cobra.Command {
	wallet_encrypt := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt an unencrypted wallet with a password",
		Run:   HandleCommandWalletEncrypt,
	}

	wallet_encrypt.Flags().String("new_password", "", "(str) password to encrypt account")
	wallet_encrypt.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return wallet_encrypt
}

func HandleCommandWalletEncrypt(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "new_password")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["new_password"]
		if exists {
			cmd.Help()
			return
		}
		params["new_password"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_encrypt", params)
}
