package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountRemove() *cobra.Command {
	account_remove := &cobra.Command{
		Use:   "remove",
		Short: "Remove an existing account.",
		Run:   HandleCommandAccountRemove,
	}

	account_remove.Flags().String("account_id", "", "(str) id of the account to remove")
	account_remove.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return account_remove
}

func HandleCommandAccountRemove(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["account_id"]
		if exists {
			cmd.Help()
			return
		}
		params["account_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_remove", params)
}
