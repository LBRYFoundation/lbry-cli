package commands_stream

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStreamList() *cobra.Command {
	stream_list := &cobra.Command{
		Use:   "list",
		Short: "List my stream claims.",
		Run:   HandleCommandStreamList,
	}

	stream_list.Flags().StringArray("name", nil, "(str or list) stream name")
	stream_list.Flags().StringArray("claim_id", nil, "(str or list) stream id")
	stream_list.Flags().Bool("is_spent", false, "(bool) shows previous stream updates and abandons")
	stream_list.Flags().String("account_id", "", "(str) id of the account to query")
	stream_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	stream_list.Flags().Int("page", -1, "(int) page to return during paginating")
	stream_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	stream_list.Flags().Bool("resolve", false, "(bool) resolves each stream to provide additional metadata")
	stream_list.Flags().Bool("no_totals", false, "(bool) do not calculate the total number of pages and items in result set (significant performance boost)")

	return stream_list
}

func HandleCommandStreamList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "resolve")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "no_totals")

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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "stream_list", params)
}
