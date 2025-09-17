package commands_purchase

import "github.com/spf13/cobra"

func CreateCommandPurchaseList() *cobra.Command {
	purchase_list := &cobra.Command{
		Use:   "list",
		Short: "List my claim purchases.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return purchase_list
}
