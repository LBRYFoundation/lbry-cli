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
	claim_type, _ := cmd.Flags().GetStringArray("claim_type")
	claim_id, _ := cmd.Flags().GetStringArray("claim_id")
	channel_id, _ := cmd.Flags().GetStringArray("channel_id")
	name, _ := cmd.Flags().GetStringArray("name")
	is_spent, _ := cmd.Flags().GetBool("is_spent")
	reposted_claim_id, _ := cmd.Flags().GetStringArray("reposted_claim_id")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	has_source, _ := cmd.Flags().GetBool("has_source")
	has_no_source, _ := cmd.Flags().GetBool("has_no_source")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")
	resolve, _ := cmd.Flags().GetBool("resolve")
	order_by, _ := cmd.Flags().GetString("order_by")
	no_totals, _ := cmd.Flags().GetBool("no_totals")
	include_received_tips, _ := cmd.Flags().GetBool("include_received_tips")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{}
	if len(claim_type) > 0 {
		params["claim_type"] = claim_type
	}
	if len(claim_id) > 0 {
		params["claim_id"] = claim_id
	}
	if len(channel_id) > 0 {
		params["channel_id"] = channel_id
	}
	if len(name) > 0 {
		params["name"] = name
	}
	if is_spent {
		params["is_spent"] = is_spent
	}
	if len(reposted_claim_id) > 0 {
		params["reposted_claim_id"] = reposted_claim_id
	}
	if account_id != "" {
		params["account_id"] = account_id
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if has_source {
		params["has_source"] = has_source
	}
	if has_no_source {
		params["has_no_source"] = has_no_source
	}
	if page != -1 {
		params["page"] = page
	}
	if page_size != -1 {
		params["page_size"] = page_size
	}
	if resolve {
		params["resolve"] = resolve
	}
	if order_by != "" {
		params["order_by"] = order_by
	}
	if no_totals {
		params["no_totals"] = no_totals
	}
	if include_received_tips {
		params["include_received_tips"] = include_received_tips
	}

	rpc.ExecuteRPCCommand("claim_list", params)
}
