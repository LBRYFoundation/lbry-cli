package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountSet() *cobra.Command {
	account_set := &cobra.Command{
		Use:   "set",
		Short: "Change various settings on an account.",
		Run:   HandleCommandAccountSet,
	}

	account_set.Flags().String("account_id", "", "(str) id of the account to change")
	account_set.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	account_set.Flags().Bool("default", false, "(bool) make this account the default")
	account_set.Flags().String("new_name", "", "(str) new name for the account")
	account_set.Flags().Int("receiving_gap", -1, "(int) set the gap for receiving addresses")
	account_set.Flags().Int("receiving_max_uses", -1, "(int) set the maximum number of times to use a receiving address")
	account_set.Flags().Int("change_gap", -1, "(int) set the gap for change addresses")
	account_set.Flags().Int("change_max_uses", -1, "(int) set the maximum number of times to use a change address")

	return account_set
}

func HandleCommandAccountSet(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "default")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "new_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "receiving_gap")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "receiving_max_uses")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "change_gap")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "change_max_uses")

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

	rpc.ExecuteRPCCommand("account_set", params)
}
