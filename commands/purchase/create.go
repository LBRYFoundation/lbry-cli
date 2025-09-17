package commands_purchase

import "github.com/spf13/cobra"

func CreateCommandPurchaseCreate() *cobra.Command {
	purchase_create := &cobra.Command{
		Use:   "create",
		Short: "Purchase a claim.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return purchase_create
}
