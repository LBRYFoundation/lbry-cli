package commands

import "github.com/spf13/cobra"

func CreateCommandAddress() *cobra.Command {
	address := &cobra.Command{
		Use:   "address",
		Short: "List, generate and verify addresses.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return address
}
