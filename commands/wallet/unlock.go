package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletUnlock() *cobra.Command {
	wallet_unlock := &cobra.Command{
		Use:   "unlock",
		Short: "Unlock an encrypted wallet",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_unlock
}
