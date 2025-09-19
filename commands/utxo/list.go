package commands_utxo

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandUTXOList() *cobra.Command {
	utxo_list := &cobra.Command{
		Use:   "list",
		Short: "List unspent transaction outputs",
		Run:   HandleCommandUTXOList,
	}

	utxo_list.Flags().String("account_id", "", "(str) id of the account to query")
	utxo_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	utxo_list.Flags().Int("page", -1, "(int) page to return during paginating")
	utxo_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return utxo_list
}

func HandleCommandUTXOList(cmd *cobra.Command, args []string) {
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")

	// Check for arguments
	if len(args) >= 1 {
		account_id = args[0]
	}

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

	rpc.ExecuteRPCCommand("utxo_list", params)
}
