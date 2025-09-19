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
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")

	// Check for arguments
	if len(args) >= 1 {
		account_id = args[0]
	}

	// Create parameter map
	params := map[string]any{}
	if account_id != "" {
		params["account_id"] = account_id
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if page != -1 {
		params["page"] = page
	}
	if page_size != -1 {
		params["page_size"] = page_size
	}

	rpc.ExecuteRPCCommand("transaction_list", params)
}
