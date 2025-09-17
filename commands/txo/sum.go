package commands_txo

import "github.com/spf13/cobra"

func CreateCommandTXOSum() *cobra.Command {
	txo_sum := &cobra.Command{
		Use:   "sum",
		Short: "Sum of transaction outputs.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return txo_sum
}
