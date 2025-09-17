package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletList() *cobra.Command {
	wallet_list := &cobra.Command{
		Use:   "list",
		Short: "List wallets.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_list
}
