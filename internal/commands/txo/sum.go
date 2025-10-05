package commands_txo

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTXOSum() *cobra.Command {
	txo_sum := &cobra.Command{
		Use:   "sum",
		Short: "Sum of transaction outputs.",
		Run:   HandleCommandTXOSum,
	}

	txo_sum.Flags().StringArray("type", nil, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_sum.Flags().StringArray("txid", nil, "(str or list) transaction id of outputs")
	txo_sum.Flags().StringArray("claim_id", nil, "(str or list) claim id")
	txo_sum.Flags().StringArray("name", nil, "(str or list) claim name")
	txo_sum.Flags().StringArray("channel_id", nil, "(str or list) claims in this channel")
	txo_sum.Flags().StringArray("not_channel_id", nil, "(str or list) claims not in this channel")
	txo_sum.Flags().Bool("is_spent", false, "(bool) only show spent txos")
	txo_sum.Flags().Bool("is_not_spent", false, "(bool) only show not spent txos")
	txo_sum.Flags().Bool("is_my_input_or_output", false, "(bool) txos which have your inputs or your outputs, if using this flag the other related flags are ignored (--is_my_output, --is_my_input, etc)")
	txo_sum.Flags().Bool("is_my_output", false, "(bool) show outputs controlled by you")
	txo_sum.Flags().Bool("is_not_my_output", false, "(bool) show outputs not controlled by you")
	txo_sum.Flags().Bool("is_my_input", false, "(bool) show outputs created by you")
	txo_sum.Flags().Bool("is_not_my_input", false, "(bool) show outputs not created by you")
	txo_sum.Flags().Bool("exclude_internal_transfers", false, "(bool) excludes any outputs that are exactly this combination: \"--is_my_input --is_my_output --type=other\" this allows to exclude \"change\" payments, this flag can be used in combination with any of the other flags")
	txo_sum.Flags().String("account_id", "", "(str) id of the account to query")
	txo_sum.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")

	return txo_sum
}

func HandleCommandTXOSum(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "type")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_no_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_input_or_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "exclude_internal_transfers")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "txo_sum", params)
}
