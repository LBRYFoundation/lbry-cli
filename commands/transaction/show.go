package commands_transaction

import "github.com/spf13/cobra"

func CreateCommandTransactionShow() *cobra.Command {
	transaction_show := &cobra.Command{
		Use:   "show",
		Short: "Get a decoded transaction from a txid",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return transaction_show
}
