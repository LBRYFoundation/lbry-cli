package commands

import commands "lbry/cli/commands/purchase"
import "github.com/spf13/cobra"

func CreateCommandPurchase() *cobra.Command {
	purchase := &cobra.Command{
		Use:   "purchase",
		Short: "List and make purchases of claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	purchase.AddCommand(commands.CreateCommandPurchaseCreate())
	purchase.AddCommand(commands.CreateCommandPurchaseList())

	return purchase
}
