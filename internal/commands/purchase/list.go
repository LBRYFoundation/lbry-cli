package commands_purchase

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPurchaseList() *cobra.Command {
	purchase_list := &cobra.Command{
		Use:   "list",
		Short: "List my claim purchases.",
		Run:   HandleCommandPurchaseList,
	}

	purchase_list.Flags().String("claim_id", "", "(str) purchases for specific claim")
	purchase_list.Flags().Bool("resolve", false, "(str) include resolved claim information")
	purchase_list.Flags().String("account_id", "", "(str) id of the account to query")
	purchase_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	purchase_list.Flags().Int("page", -1, "(int) page to return during paginating")
	purchase_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return purchase_list
}

func HandleCommandPurchaseList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "resolve")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["claim_id"]
		if exists {
			cmd.Help()
			return
		}
		params["claim_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "purchase_list", params)
}
