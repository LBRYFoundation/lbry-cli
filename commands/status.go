package commands

import "github.com/spf13/cobra"

func CreateCommandStatus() *cobra.Command {
	status := &cobra.Command{
		Use:   "status",
		Short: "Get daemon status",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return status
}
