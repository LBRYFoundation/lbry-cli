package commands_address

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAddressUnused() *cobra.Command {
	address_unused := &cobra.Command{
		Use:   "unused",
		Short: "Return an address containing no balance, will create a new address if there is none.",
		Run:   HandleCommandAddressUnused,
	}

	address_unused.Flags().String("account_id", "", "(str) id of the account to use")
	address_unused.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return address_unused
}

func HandleCommandAddressUnused(cmd *cobra.Command, args []string) {
	account_id, _ := cmd.Flags().GetString("account_id")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{}
	if account_id != "" {
		params["account_id"] = account_id
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}

	rpc.ExecuteRPCCommand("address_unused", params)
}
