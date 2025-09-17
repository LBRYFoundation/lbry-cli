package commands_preference

import "github.com/spf13/cobra"

func CreateCommandPreferenceGet() *cobra.Command {
	preference_get := &cobra.Command{
		Use:   "get",
		Short: "Get preference value for key or all values if not key is passed in.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return preference_get
}
