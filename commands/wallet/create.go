package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletCreate() *cobra.Command {
	wallet_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_create
}
