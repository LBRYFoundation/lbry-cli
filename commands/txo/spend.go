package commands_txo

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTXOSpend() *cobra.Command {
	txo_spend := &cobra.Command{
		Use:   "spend",
		Short: "Spend transaction outputs, batching into multiple transactions as necessary.",
		Run:   HandleCommandTXOSpend,
	}

	txo_spend.Flags().StringArray("type", nil, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_spend.Flags().StringArray("txid", nil, "(str or list) transaction id of outputs")
	txo_spend.Flags().StringArray("claim_id", nil, "(str or list) claim id")
	txo_spend.Flags().StringArray("channel_id", nil, "(str or list) claims in this channel")
	txo_spend.Flags().StringArray("not_channel_id", nil, "(str or list) claims not in this channel")
	txo_spend.Flags().StringArray("name", nil, "(str or list) claim name")
	txo_spend.Flags().Bool("is_my_input", false, "(bool) show outputs created by you")
	txo_spend.Flags().Bool("is_not_my_input", false, "(bool) show outputs not created by you")
	txo_spend.Flags().Bool("exclude_internal_transfers", false, "(bool) excludes any outputs that are exactly this combination: \"--is_my_input --is_my_output --type=other\" this allows to exclude \"change\" payments, this flag can be used in combination with any of the other flags")
	txo_spend.Flags().String("account_id", "", "(str) id of the account to query")
	txo_spend.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	txo_spend.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	txo_spend.Flags().Bool("blocking", false, "(bool) wait until abandon is in mempool")
	txo_spend.Flags().Int("batch_size", -1, "(int) number of txos to spend per transactions")
	txo_spend.Flags().Bool("include_full_tx", false, "(bool) include entire tx in output and not just the txid")

	return txo_spend
}

func HandleCommandTXOSpend(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "type")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_not_my_input")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "exclude_internal_transfers")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "batch_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_full_tx")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "txo_spend", params)
}
