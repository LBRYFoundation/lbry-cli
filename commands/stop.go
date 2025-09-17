package commands

import "github.com/spf13/cobra"

func CreateCommandStop() *cobra.Command {
	stop := &cobra.Command{
		Use:   "stop",
		Short: "Stop lbrynet API server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stop
}
