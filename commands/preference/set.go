package commands_preference

import "github.com/spf13/cobra"

func CreateCommandPreferenceSet() *cobra.Command {
	preference_set := &cobra.Command{
		Use:   "set",
		Short: "Set preferences",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return preference_set
}
