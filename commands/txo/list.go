package commands_txo

import "github.com/spf13/cobra"

func CreateCommandTXOList() *cobra.Command {
	txo_list := &cobra.Command{
		Use:   "list",
		Short: "List my transaction outputs.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return txo_list
}
