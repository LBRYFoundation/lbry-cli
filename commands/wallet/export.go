package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletExport() *cobra.Command {
	wallet_export := &cobra.Command{
		Use:   "export",
		Short: "Exports encrypted wallet data if password is supplied; otherwise plain JSON.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_export
}
