package commands_preference

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPreferenceGet() *cobra.Command {
	preference_get := &cobra.Command{
		Use:   "get",
		Short: "Get preference value for key or all values if not key is passed in.",
		Run:   HandleCommandPreferenceGet,
	}

	preference_get.Flags().String("key", "", "(str) key associated with value")
	preference_get.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return preference_get
}

func HandleCommandPreferenceGet(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "key")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "preference_get", params)
}
