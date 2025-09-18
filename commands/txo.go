package commands

import (
	commands "lbry/cli/commands/txo"

	"github.com/spf13/cobra"
)

func CreateCommandTXO() *cobra.Command {
	txo := &cobra.Command{
		Use:   "txo",
		Short: "List and sum transaction outputs.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	txo.AddCommand(commands.CreateCommandTXOList())
	txo.AddCommand(commands.CreateCommandTXOPlot())
	txo.AddCommand(commands.CreateCommandTXOSpend())
	txo.AddCommand(commands.CreateCommandTXOSum())

	return txo
}
