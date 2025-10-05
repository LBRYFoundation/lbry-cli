package commands

import (
	commands "lbry/cli/commands/settings"

	"github.com/spf13/cobra"
)

func CreateCommandSettings() *cobra.Command {
	settings := &cobra.Command{
		Use:   "settings",
		Short: "Settings management.",
	}

	settings.AddCommand(commands.CreateCommandSettingsClear())
	settings.AddCommand(commands.CreateCommandSettingsGet())
	settings.AddCommand(commands.CreateCommandSettingsSet())

	return settings
}
