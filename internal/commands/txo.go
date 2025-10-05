package commands

import (
	commands "lbry/cli/internal/commands/txo"

	"github.com/spf13/cobra"
)

func CreateCommandTXO() *cobra.Command {
	txo := &cobra.Command{
		Use:   "txo",
		Short: "List and sum transaction outputs.",
	}

	txo.AddCommand(commands.CreateCommandTXOList())
	txo.AddCommand(commands.CreateCommandTXOPlot())
	txo.AddCommand(commands.CreateCommandTXOSpend())
	txo.AddCommand(commands.CreateCommandTXOSum())

	return txo
}
