package commands

import "github.com/spf13/cobra"

func CreateCommandSync() *cobra.Command {
	sync := &cobra.Command{
		Use:   "sync",
		Short: "Wallet synchronization.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return sync
}
