package commands_settings

import "github.com/spf13/cobra"

func CreateCommandSettingsSet() *cobra.Command {
	settings_set := &cobra.Command{
		Use:   "set",
		Short: "Set daemon settings",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return settings_set
}
