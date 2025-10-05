package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountList() *cobra.Command {
	account_list := &cobra.Command{
		Use:   "list",
		Short: "List details of all of the accounts or a specific account.",
		Run:   HandleCommandAccountList,
	}

	//account_list.Flags().String("account_id", "", "(str) If provided only the balance for this account will be given")
	account_list.Flags().String("wallet_id", "", "(str) accounts in specific wallet")
	account_list.Flags().Int("confirmations", -1, "(int) required confirmations (default: 0)")
	account_list.Flags().Bool("include_claims", false, "(bool) include claims, requires than a LBC account is specified (default: false)")
	account_list.Flags().Bool("show_seed", false, "(bool) show the seed for the account")
	account_list.Flags().Int("page", -1, "(int) page to return during paginating")
	account_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return account_list
}

func HandleCommandAccountList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "confirmations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_claims")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "show_seed")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_list", params)
}
