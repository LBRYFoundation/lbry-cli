package commands_address

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAddressList() *cobra.Command {
	address_list := &cobra.Command{
		Use:   "list",
		Short: "List account addresses or details of single address.",
		Run:   HandleCommandAddressList,
	}

	address_list.Flags().String("address", "", "(str) just show details for single address")
	address_list.Flags().String("account_id", "", "(str) id of the account to use")
	address_list.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	address_list.Flags().String("page", "", "(int) page to return during paginating")
	address_list.Flags().String("page_size", "", "(int) number of items on page during pagination")

	return address_list
}

func HandleCommandAddressList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "address_list", params)
}
