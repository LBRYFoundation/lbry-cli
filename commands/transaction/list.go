package commands_transaction

import "github.com/spf13/cobra"

func CreateCommandTransactionList() *cobra.Command {
	transaction_list := &cobra.Command{
		Use:   "list",
		Short: "List transactions belonging to wallet",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return transaction_list
}
