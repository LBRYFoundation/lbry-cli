package commands_txo

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTXOPlot() *cobra.Command {
	txo_plot := &cobra.Command{
		Use:   "plot",
		Short: "Plot transaction output sum over days.",
		Run:   HandleCommandTXOPlot,
	}

	txo_plot.Flags().StringArray("type", []string{}, "(str or list) claim type: stream, channel, support, purchase, collection, repost, other")
	txo_plot.Flags().StringArray("txid", []string{}, "(str or list) transaction id of outputs")
	txo_plot.Flags().StringArray("claim_id", []string{}, "(str or list) claim id")
	txo_plot.Flags().StringArray("name", []string{}, "(str or list) claim name")
	txo_plot.Flags().StringArray("channel_id", []string{}, "(str or list) claims in this channel")
	txo_plot.Flags().StringArray("not_channel_id", []string{}, "(str or list) claims not in this channel")
	txo_plot.Flags().Bool("is_spent", false, "(bool) only show spent txos")
	txo_plot.Flags().Bool("is_not_spent", false, "(bool) only show not spent txos")
	txo_plot.Flags().Bool("is_my_input_or_output", false, "(bool) txos which have your inputs or your outputs, if using this flag the other related flags are ignored (--is_my_output, --is_my_input, etc)")
	txo_plot.Flags().Bool("is_my_output", false, "(bool) show outputs controlled by you")
	txo_plot.Flags().Bool("is_not_my_output", false, "(bool) show outputs not controlled by you")
	txo_plot.Flags().Bool("is_my_input", false, "(bool) show outputs created by you")
	txo_plot.Flags().Bool("is_not_my_input", false, "(bool) show outputs not created by you")
	txo_plot.Flags().Bool("exclude_internal_transfers", false, "(bool) excludes any outputs that are exactly this combination: \"--is_my_input --is_my_output --type=other\" this allows to exclude \"change\" payments, this flag can be used in combination with any of the other flags")
	txo_plot.Flags().String("account_id", "", "(str) id of the account to query")
	txo_plot.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	txo_plot.Flags().Int("days_back", -1, "(int) number of days back from today (not compatible with --start_day, --days_after, --end_day)")
	txo_plot.Flags().String("start_day", "", "(date) start on specific date (YYYY-MM-DD) (instead of --days_back)")
	txo_plot.Flags().Int("days_after", -1, "(int) end number of days after --start_day (instead of --end_day)")
	txo_plot.Flags().String("end_day", "", "(date) end on specific date (YYYY-MM-DD) (instead of --days_after)")

	return txo_plot
}

func HandleCommandTXOPlot(cmd *cobra.Command, args []string) {
	_type, _ := cmd.Flags().GetStringArray("type")
	txid, _ := cmd.Flags().GetStringArray("txid")
	claim_id, _ := cmd.Flags().GetStringArray("claim_id")
	name, _ := cmd.Flags().GetStringArray("name")
	channel_id, _ := cmd.Flags().GetStringArray("channel_id")
	not_channel_id, _ := cmd.Flags().GetStringArray("not_channel_id")
	is_spent, _ := cmd.Flags().GetBool("is_spent")
	is_not_spent, _ := cmd.Flags().GetBool("is_not_spent")
	is_my_input_or_output, _ := cmd.Flags().GetBool("is_my_input_or_output")
	is_my_output, _ := cmd.Flags().GetBool("is_my_output")
	is_not_my_output, _ := cmd.Flags().GetBool("is_not_my_output")
	is_my_input, _ := cmd.Flags().GetBool("is_my_input")
	is_not_my_input, _ := cmd.Flags().GetBool("is_not_my_input")
	exclude_internal_transfers, _ := cmd.Flags().GetBool("exclude_internal_transfers")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	days_back, _ := cmd.Flags().GetInt("days_back")
	start_day, _ := cmd.Flags().GetString("start_day")
	days_after, _ := cmd.Flags().GetInt("days_after")
	end_day, _ := cmd.Flags().GetString("end_day")

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
	if len(name) > 0 {
		params["name"] = name
	}
	if len(channel_id) > 0 {
		params["channel_id"] = channel_id
	}
	if len(not_channel_id) > 0 {
		params["not_channel_id"] = not_channel_id
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
	if account_id != "" {
		params["account_id"] = account_id
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}

	if days_back != -1 {
		params["days_back"] = days_back
	}
	if start_day != "" {
		params["start_day"] = start_day
	}
	if days_after != -1 {
		params["days_after"] = days_after
	}
	if end_day != "" {
		params["end_day"] = end_day
	}

	rpc.ExecuteRPCCommand("txo_plot", params)
}
