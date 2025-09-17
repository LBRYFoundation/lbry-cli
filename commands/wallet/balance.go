package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletBalance() *cobra.Command {
	wallet_balance := &cobra.Command{
		Use:   "balance",
		Short: "Return the balance of a wallet",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_balance
}
