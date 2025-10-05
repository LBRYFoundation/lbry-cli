package commands

import (
	commands "lbry/cli/internal/commands/utxo"

	"github.com/spf13/cobra"
)

func CreateCommandUTXO() *cobra.Command {
	utxo := &cobra.Command{
		Use:   "utxo",
		Short: "Unspent transaction management.",
	}

	utxo.AddCommand(commands.CreateCommandUTXOList())
	utxo.AddCommand(commands.CreateCommandUTXORelease())

	return utxo
}
