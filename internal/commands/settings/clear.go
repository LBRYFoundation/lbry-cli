package commands_settings

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSettingsClear() *cobra.Command {
	settings_clear := &cobra.Command{
		Use:   "clear",
		Short: "Clear daemon settings",
		Run:   HandleCommandSettingsClear,
	}

	return settings_clear
}

func HandleCommandSettingsClear(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["key"]
		if exists {
			cmd.Help()
			return
		}
		params["key"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, exists := params["key"]
	if !exists {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "settings_clear", params)
}
