package commands

import commands "lbry/cli/commands/utxo"
import "github.com/spf13/cobra"

func CreateCommandUTXO() *cobra.Command {
	utxo := &cobra.Command{
		Use:   "utxo",
		Short: "Unspent transaction management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	utxo.AddCommand(commands.CreateCommandUTXOList())
	utxo.AddCommand(commands.CreateCommandUTXORelease())

	return utxo
}
