package commands_utxo

import "github.com/spf13/cobra"

func CreateCommandUTXOList() *cobra.Command {
	utxo_list := &cobra.Command{
		Use:   "list",
		Short: "List unspent transaction outputs",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return utxo_list
}
