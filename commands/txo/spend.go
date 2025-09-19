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

	txo_spend.Flags().StringArray("type", []string{}, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_spend.Flags().StringArray("txid", []string{}, "(str or list) transaction id of outputs")
	txo_spend.Flags().StringArray("claim_id", []string{}, "(str or list) claim id")
	txo_spend.Flags().StringArray("channel_id", []string{}, "(str or list) claims in this channel")
	txo_spend.Flags().StringArray("not_channel_id", []string{}, "(str or list) claims not in this channel")
	txo_spend.Flags().StringArray("name", []string{}, "(str or list) claim name")
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
	_type, _ := cmd.Flags().GetStringArray("type")
	txid, _ := cmd.Flags().GetStringArray("txid")
	claim_id, _ := cmd.Flags().GetStringArray("claim_id")
	channel_id, _ := cmd.Flags().GetStringArray("channel_id")
	not_channel_id, _ := cmd.Flags().GetStringArray("not_channel_id")
	name, _ := cmd.Flags().GetStringArray("name")
	is_my_input, _ := cmd.Flags().GetBool("is_my_input")
	is_not_my_input, _ := cmd.Flags().GetBool("is_not_my_input")
	exclude_internal_transfers, _ := cmd.Flags().GetBool("exclude_internal_transfers")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	preview, _ := cmd.Flags().GetBool("preview")
	blocking, _ := cmd.Flags().GetBool("blocking")
	batch_size, _ := cmd.Flags().GetInt("batch_size")
	include_full_tx, _ := cmd.Flags().GetBool("include_full_tx")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	// Create parameter map
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
	if is_my_input {
		params["is_my_input"] = is_my_input
	}
	if is_not_my_input {
		params["is_not_my_input"] = is_not_my_input
	}
	if exclude_internal_transfers {
		params["exclude_internal_transfers"] = exclude_internal_transfers
	}
	if account_id != "" {
		params["account_id"] = account_id
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if preview {
		params["preview"] = preview
	}
	if blocking {
		params["blocking"] = blocking
	}
	if batch_size != -1 {
		params["batch_size"] = batch_size
	}
	if include_full_tx {
		params["include_full_tx"] = include_full_tx
	}

	rpc.ExecuteRPCCommand("txo_spend", params)
}
