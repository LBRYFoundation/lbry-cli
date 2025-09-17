package commands_txo

import "github.com/spf13/cobra"

func CreateCommandTXOSpend() *cobra.Command {
	txo_spend := &cobra.Command{
		Use:   "spend",
		Short: "Spend transaction outputs, batching into multiple transactions as necessary.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return txo_spend
}
