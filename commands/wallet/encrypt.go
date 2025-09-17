package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletEncrypt() *cobra.Command {
	wallet_encrypt := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt an unencrypted wallet with a password",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_encrypt
}
