package commands_settings

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSettingsGet() *cobra.Command {
	settings_get := &cobra.Command{
		Use:   "get",
		Short: "Get daemon settings",
		Run:   HandleCommandSettingsGet,
	}

	return settings_get
}

func HandleCommandSettingsGet(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("settings_get")
}
