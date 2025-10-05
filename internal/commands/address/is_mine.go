package commands_address

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAddressIsMine() *cobra.Command {
	address_is_mine := &cobra.Command{
		Use:   "is_mine",
		Short: "Checks if an address is associated with the current wallet.",
		Run:   HandleCommandAddressIsMine,
	}

	address_is_mine.Flags().String("address", "", "(str) address to check")
	address_is_mine.Flags().String("account_id", "", "(str) id of the account to use")
	address_is_mine.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return address_is_mine
}

func HandleCommandAddressIsMine(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["address"]
		if exists {
			cmd.Help()
			return
		}
		params["address"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["account_id"]
		if exists {
			cmd.Help()
			return
		}
		params["account_id"] = args[1]
	}
	if len(args) > 2 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "address_is_mine", params)
}
