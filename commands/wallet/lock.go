package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletLock() *cobra.Command {
	wallet_lock := &cobra.Command{
		Use:   "lock",
		Short: "Lock an unlocked wallet",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_lock
}
