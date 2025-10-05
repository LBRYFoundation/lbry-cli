package commands_claim

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandClaimList() *cobra.Command {
	claim_list := &cobra.Command{
		Use:   "list",
		Short: "List my stream and channel claims.",
		Run:   HandleCommandClaimList,
	}

	claim_list.Flags().StringArray("claim_type", nil, "(str or list) claim id")
	claim_list.Flags().StringArray("claim_id", nil, "(str or list) claim id")
	claim_list.Flags().StringArray("channel_id", nil, "(str or list) claims in this channel")
	claim_list.Flags().StringArray("name", nil, "(str or list) claim name")
	claim_list.Flags().Bool("is_spent", false, "(bool) only show spent txos")
	claim_list.Flags().StringArray("reposted_claim_id", nil, "(str or list) reposted claim id")
	claim_list.Flags().String("account_id", "", "(str) id of the account to query")
	claim_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	claim_list.Flags().Bool("has_source", false, "(bool) list claims containing a source field")
	claim_list.Flags().Bool("has_no_source", false, "(bool) list claims not containing a source field")
	claim_list.Flags().Int("page", -1, "(int) page to return during paginating")
	claim_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	claim_list.Flags().Bool("resolve", false, "(bool) resolves each claim to provide additional metadata")
	claim_list.Flags().String("order_by", "", "(str) field to order by: 'name', 'height', 'amount'")
	claim_list.Flags().Bool("no_totals", false, "(bool) do not calculate the total number of pages and items in result set (significant performance boost)")
	claim_list.Flags().Bool("include_received_tips", false, "(bool) calculate the amount of tips received for claim outputs")

	return claim_list
}

func HandleCommandClaimList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_type")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "reposted_claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "has_source")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "has_no_source")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "resolve")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "order_by")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "no_totals")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_received_tips")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "claim_list", params)
}
