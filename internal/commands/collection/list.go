package commands_collection

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandCollectionList() *cobra.Command {
	collection_list := &cobra.Command{
		Use:   "list",
		Short: "List my collection claims.",
		Run:   HandleCommandCollectionList,
	}

	collection_list.Flags().Bool("resolve", false, "(bool) resolve collection claim")
	collection_list.Flags().Int("resolve_claims", -1, "(int) resolve every claim")
	collection_list.Flags().String("account_id", "", "(str) id of the account to use")
	collection_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	collection_list.Flags().Int("page", -1, "(int) page to return during paginating")
	collection_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return collection_list
}

func HandleCommandCollectionList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "resolve")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "resolve_claims")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "collection_list", params)
}
