package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletReconnect() *cobra.Command {
	wallet_reconnect := &cobra.Command{
		Use:   "reconnect",
		Short: "Reconnects ledger network client, applying new configurations.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_reconnect
}
