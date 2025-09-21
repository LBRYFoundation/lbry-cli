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

	rpc.ExecuteRPCCommand("utxo_list", params)
}
