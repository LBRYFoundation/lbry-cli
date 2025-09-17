package commands

import "github.com/spf13/cobra"

func CreateCommandStart() *cobra.Command {
	start := &cobra.Command{
		Use:   "start",
		Short: "Start LBRY Network interface.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return start
}
