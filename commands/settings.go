package commands

import "github.com/spf13/cobra"

func CreateCommandSettings() *cobra.Command {
	settings := &cobra.Command{
		Use:   "settings",
		Short: "Settings management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return settings
}
