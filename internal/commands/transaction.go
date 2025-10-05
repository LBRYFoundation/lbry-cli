package commands

import (
	commands "lbry/cli/internal/commands/transaction"

	"github.com/spf13/cobra"
)

func CreateCommandTransaction() *cobra.Command {
	transaction := &cobra.Command{
		Use:   "transaction",
		Short: "Transaction management.",
	}

	transaction.AddCommand(commands.CreateCommandTransactionList())
	transaction.AddCommand(commands.CreateCommandTransactionShow())

	return transaction
}
