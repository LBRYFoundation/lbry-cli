package commands_txo

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTXOList() *cobra.Command {
	txo_list := &cobra.Command{
		Use:   "list",
		Short: "List my transaction outputs.",
		Run:   HandleCommandTXOList,
	}

	txo_list.Flags().StringArray("type", []string{}, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_list.Flags().StringArray("txid", []string{}, "(str or list) transaction id of outputs")
	txo_list.Flags().StringArray("claim_id", []string{}, "(str or list) claim id")
	txo_list.Flags().StringArray("channel_id", []string{}, "(str or list) claims in this channel")
	txo_list.Flags().StringArray("not_channel_id", []string{}, "(str or list) claims not in this channel")
	txo_list.Flags().StringArray("name", []string{}, "(str or list) claim name")
	txo_list.Flags().Bool("is_spent", false, "(bool) only show spent txos")
	txo_list.Flags().Bool("is_not_spent", false, "(bool) only show not spent txos")
	txo_list.Flags().Bool("is_my_input_or_output", false, "(bool) txos which have your inputs or your outputs, if using this flag the other related flags are ignored (--is_my_output, --is_my_input, etc)")
	txo_list.Flags().Bool("is_my_output", false, "(bool) show outputs controlled by you")
	txo_list.Flags().Bool("is_not_my_output", false, "(bool) show outputs not controlled by you")
	txo_list.Flags().Bool("is_my_input", false, "(bool) show outputs created by you")
	txo_list.Flags().Bool("is_not_my_input", false, "(bool) show outputs not created by you")
	txo_list.Flags().Bool("exclude_internal_transfers", false, "(bool) excludes any outputs that are exactly this combination: \"--is_my_input --is_my_output --type=other\" this allows to exclude \"change\" payments, this flag can be used in combination with any of the other flags")
	txo_list.Flags().Bool("include_received_tips", false, "(bool) calculate the amount of tips received for claim outputs")
	txo_list.Flags().String("account_id", "", "(str) id of the account to query")
	txo_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	txo_list.Flags().Int("page", -1, "(int) page to return during paginating")
	txo_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	txo_list.Flags().Bool("resolve", false, "(bool) resolves each claim to provide additional metadata")
	txo_list.Flags().String("order_by", "", "(str) field to order by: 'name', 'height', 'amount' and 'none'")
	txo_list.Flags().Bool("no_totals", false, "(bool) do not calculate the total number of pages and items in result set (significant performance boost)")

	return txo_list
}

func HandleCommandTXOList(cmd *cobra.Command, args []string) {
	_type, _ := cmd.Flags().GetStringArray("type")
	txid, _ := cmd.Flags().GetStringArray("txid")
	claim_id, _ := cmd.Flags().GetStringArray("claim_id")
	channel_id, _ := cmd.Flags().GetStringArray("channel_id")
	not_channel_id, _ := cmd.Flags().GetStringArray("not_channel_id")
	name, _ := cmd.Flags().GetStringArray("name")
	is_spent, _ := cmd.Flags().GetBool("is_spent")
	is_not_spent, _ := cmd.Flags().GetBool("is_not_spent")
	is_my_input_or_output, _ := cmd.Flags().GetBool("is_my_input_or_output")
	is_my_output, _ := cmd.Flags().GetBool("is_my_output")
	is_not_my_output, _ := cmd.Flags().GetBool("is_not_my_output")
	is_my_input, _ := cmd.Flags().GetBool("is_my_input")
	is_not_my_input, _ := cmd.Flags().GetBool("is_not_my_input")
	exclude_internal_transfers, _ := cmd.Flags().GetBool("exclude_internal_transfers")
	include_received_tips, _ := cmd.Flags().GetBool("include_received_tips")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")
	resolve, _ := cmd.Flags().GetBool("resolve")
	order_by, _ := cmd.Flags().GetString("order_by")
	no_totals, _ := cmd.Flags().GetBool("no_totals")

	// Check for arguments
	if len(args) != 0 {
		cmd.Help()
		return
	}

	params := map[string]any{}
	if len(_type) > 0 {
		params["type"] = _type
	}
	if len(txid) > 0 {
		params["txid"] = txid
	}
	if len(claim_id) > 0 {
		params["claim_id"] = claim_id
	}
	if len(channel_id) > 0 {
		params["channel_id"] = channel_id
	}
	if len(not_channel_id) > 0 {
		params["not_channel_id"] = not_channel_id
	}
	if len(name) > 0 {
		params["name"] = name
	}
	if is_spent {
		params["is_spent"] = is_spent
	}
	if is_not_spent {
		params["is_not_spent"] = is_not_spent
	}
	if is_my_input_or_output {
		params["is_my_input_or_output"] = is_my_input_or_output
	}
	if is_my_output {
		params["is_my_output"] = is_my_output
	}
	if is_not_my_output {
		params["is_not_my_output"] = is_not_my_output
	}
	if is_my_input {
		params["is_my_input"] = is_my_input
	}
	if is_not_my_input {
		params["is_not_my_input"] = is_not_my_input
	}
	if exclude_internal_transfers {
		params["exclude_internal_transfers"] = exclude_internal_transfers
	}
	if include_received_tips {
		params["include_received_tips"] = include_received_tips
	}
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
	if resolve {
		params["resolve"] = resolve
	}
	if order_by != "" {
		params["order_by"] = order_by
	}
	if no_totals {
		params["no_totals"] = no_totals
	}

	rpc.ExecuteRPCCommand("txo_list", params)
}
