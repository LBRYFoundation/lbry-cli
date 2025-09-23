package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountCreate() *cobra.Command {
	account_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new account. Specify --single_key if you want to use the same address for all transactions (not recommended).",
		Run:   HandleCommandAccountCreate,
	}

	account_create.Flags().String("account_name", "", "(str) name of the account to create")
	account_create.Flags().Bool("single_key", false, "(bool) create single key account, default is multi-key")
	account_create.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return account_create
}

func HandleCommandAccountCreate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "single_key")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["account_name"]
		if exists {
			cmd.Help()
			return
		}
		params["account_name"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_create", params)
}
