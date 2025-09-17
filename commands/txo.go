package commands

import "github.com/spf13/cobra"

func CreateCommandTXO() *cobra.Command {
	txo := &cobra.Command{
		Use:   "txo",
		Short: "List and sum transaction outputs.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return txo
}
