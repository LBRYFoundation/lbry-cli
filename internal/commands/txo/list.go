package commands_txo

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTXOList() *cobra.Command {
	txo_list := &cobra.Command{
		Use:   "list",
		Short: "List my transaction outputs.",
		Run:   HandleCommandTXOList,
	}

	txo_list.Flags().StringArray("type", nil, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_list.Flags().StringArray("txid", nil, "(str or list) transaction id of outputs")
	txo_list.Flags().StringArray("claim_id", nil, "(str or list) claim id")
	txo_list.Flags().StringArray("channel_id", nil, "(str or list) claims in this channel")
	txo_list.Flags().StringArray("not_channel_id", nil, "(str or list) claims not in this channel")
	txo_list.Flags().StringArray("name", nil, "(str or list) claim name")
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
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "type")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_input_or_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "exclude_internal_transfers")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_received_tips")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "resolve")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "order_by")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "no_totals")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "txo_list", params)
}
