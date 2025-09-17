package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletDecrypt() *cobra.Command {
	wallet_decrypt := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt an encrypted wallet, this will remove the wallet password. The wallet must be unlocked to decrypt it",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_decrypt
}
