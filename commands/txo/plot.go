package commands_txo

import "github.com/spf13/cobra"

func CreateCommandTXOPlot() *cobra.Command {
	txo_plot := &cobra.Command{
		Use:   "plot",
		Short: "Plot transaction output sum over days.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return txo_plot
}
