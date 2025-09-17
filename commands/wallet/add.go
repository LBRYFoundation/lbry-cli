package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletAdd() *cobra.Command {
	wallet_add := &cobra.Command{
		Use:   "add",
		Short: "Add existing wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_add
}
