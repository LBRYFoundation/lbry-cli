package commands

import (
	commands "lbry/cli/internal/commands/preference"

	"github.com/spf13/cobra"
)

func CreateCommandPreference() *cobra.Command {
	preference := &cobra.Command{
		Use:   "preference",
		Short: "Preferences management.",
	}

	preference.AddCommand(commands.CreateCommandPreferenceGet())
	preference.AddCommand(commands.CreateCommandPreferenceSet())

	return preference
}
