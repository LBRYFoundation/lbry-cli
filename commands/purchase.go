package commands

import (
	commands "lbry/cli/commands/purchase"

	"github.com/spf13/cobra"
)

func CreateCommandPurchase() *cobra.Command {
	purchase := &cobra.Command{
		Use:   "purchase",
		Short: "List and make purchases of claims.",
	}

	purchase.AddCommand(commands.CreateCommandPurchaseCreate())
	purchase.AddCommand(commands.CreateCommandPurchaseList())

	return purchase
}
