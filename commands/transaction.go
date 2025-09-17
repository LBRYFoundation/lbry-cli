package commands

import commands "lbry/cli/commands/transaction"
import "github.com/spf13/cobra"

func CreateCommandTransaction() *cobra.Command {
	transaction := &cobra.Command{
		Use:   "transaction",
		Short: "Transaction management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	transaction.AddCommand(commands.CreateCommandTransactionList())
	transaction.AddCommand(commands.CreateCommandTransactionShow())

	return transaction
}
