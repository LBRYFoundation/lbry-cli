package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletRemove() *cobra.Command {
	wallet_remove := &cobra.Command{
		Use:   "remove",
		Short: "Remove an existing wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_remove
}
