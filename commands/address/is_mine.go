package commands_address

import (
	"lbry/cli/rpc"

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
	address, _ := cmd.Flags().GetString("address")
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		address = args[0]
	}
	if len(args) >= 2 {
		account_id = args[1]
	}
	if len(args) > 2 {
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

	rpc.ExecuteRPCCommand("address_is_mine", params)
}
