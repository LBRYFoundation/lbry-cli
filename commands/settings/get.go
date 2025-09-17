package commands_settings

import "github.com/spf13/cobra"

func CreateCommandSettingsGet() *cobra.Command {
	settings_get := &cobra.Command{
		Use:   "get",
		Short: "Get daemon settings",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return settings_get
}
