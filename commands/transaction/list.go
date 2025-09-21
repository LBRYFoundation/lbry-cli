package commands_transaction

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTransactionList() *cobra.Command {
	transaction_list := &cobra.Command{
		Use:   "list",
		Short: "List transactions belonging to wallet",
		Run:   HandleCommandTransactionList,
	}

	transaction_list.Flags().String("account_id", "", "(str) id of the account to query")
	transaction_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	transaction_list.Flags().Int("page", -1, "(int) page to return during paginating")
	transaction_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return transaction_list
}

func HandleCommandTransactionList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

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

	rpc.ExecuteRPCCommand("transaction_list", params)
}
