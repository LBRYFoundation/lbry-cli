package commands_settings

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSettingsSet() *cobra.Command {
	settings_set := &cobra.Command{
		Use:   "set",
		Short: "Set daemon settings",
		Run:   HandleCommandSettingsSet,
	}

	return settings_set
}

func HandleCommandSettingsSet(cmd *cobra.Command, args []string) {
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
	if len(args) >= 2 {
		_, exists := params["value"]
		if exists {
			cmd.Help()
			return
		}
		params["value"] = args[1]
	}
	if len(args) > 2 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, existsKey := params["key"]
	if !existsKey {
		cmd.Help()
		return
	}
	// Check for required parameters
	_, existsValue := params["value"]
	if !existsValue {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("settings_set", params)
}
