package commands

import "github.com/spf13/cobra"

func CreateCommandWallet() *cobra.Command {
	wallet := &cobra.Command{
		Use:   "wallet",
		Short: "Create, modify and inspect wallets.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet
}
