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
	address, _ := cmd.Flags().GetString("address")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{}
	if address != "" {
		params["address"] = address
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

	rpc.ExecuteRPCCommand("address_list", params)
}
