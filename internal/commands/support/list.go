package commands_support

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSupportList() *cobra.Command {
	support_list := &cobra.Command{
		Use:   "list",
		Short: "List staked supports and sent/received tips.",
		Run:   HandleCommandSupportList,
	}

	support_list.Flags().StringArray("name", nil, "(str or list) claim name")
	support_list.Flags().StringArray("claim_id", nil, "(str or list) claim id")
	support_list.Flags().Bool("received", false, "(bool) only show received (tips)")
	support_list.Flags().Bool("sent", false, "(bool) only show sent (tips)")
	support_list.Flags().Bool("staked", false, "(bool) only show my staked supports")
	support_list.Flags().Bool("is_spent", false, "(bool) show abandoned supports")
	support_list.Flags().String("account_id", "", "(str) id of the account to query")
	support_list.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	support_list.Flags().Int("page", -1, "(int) page to return during paginating")
	support_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	support_list.Flags().Bool("no_totals", false, "(bool) do not calculate the total number of pages and items in result set (significant performance boost)")

	return support_list
}

func HandleCommandSupportList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "received")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "sent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "staked")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_spent")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "support_list", params)
}
