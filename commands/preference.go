package commands

import commands "lbry/cli/commands/preference"
import "github.com/spf13/cobra"

func CreateCommandPreference() *cobra.Command {
	preference := &cobra.Command{
		Use:   "preference",
		Short: "Preferences management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	preference.AddCommand(commands.CreateCommandPreferenceGet())
	preference.AddCommand(commands.CreateCommandPreferenceSet())

	return preference
}
