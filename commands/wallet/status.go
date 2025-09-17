package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletStatus() *cobra.Command {
	wallet_status := &cobra.Command{
		Use:   "status",
		Short: "Status of wallet including encryption/lock state.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_status
}
