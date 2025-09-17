package commands_settings

import "github.com/spf13/cobra"

func CreateCommandSettingsClear() *cobra.Command {
	settings_clear := &cobra.Command{
		Use:   "clear",
		Short: "Clear daemon settings",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return settings_clear
}
