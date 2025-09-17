package commands

import "github.com/spf13/cobra"

func CreateCommandPreference() *cobra.Command {
	preference := &cobra.Command{
		Use:   "preference",
		Short: "Preferences management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return preference
}
