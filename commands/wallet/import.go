package commands_wallet

import "github.com/spf13/cobra"

func CreateCommandWalletImport() *cobra.Command {
	wallet_import := &cobra.Command{
		Use:   "import",
		Short: "Import wallet data and merge accounts and preferences. Data is expected to be JSON if password is not supplied.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return wallet_import
}
