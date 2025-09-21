package commands_preference

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPreferenceSet() *cobra.Command {
	preference_set := &cobra.Command{
		Use:   "set",
		Short: "Set preferences",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	preference_set.Flags().String("key", "", "(str) key associated with value")
	preference_set.Flags().String("value", "", "(str) key associated with value")
	preference_set.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return preference_set
}

func HandleCommandPreferenceSet(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "key")
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "value")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

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
	_, existsValue := params["value"]
	if !existsValue {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("preference_set", params)
}
