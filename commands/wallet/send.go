package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletSend() *cobra.Command {
	wallet_send := &cobra.Command{
		Use:   "send",
		Short: "Send the same number of credits to multiple addresses using all accounts in wallet to fund the transaction and the default account to receive any change.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_send
}
