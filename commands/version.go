package commands

import "github.com/spf13/cobra"

func CreateCommandVersion() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Get lbrynet API server version information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return version
}
