package commands_address

import (
	"lbry/cli/internal/rpc"

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
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "address_unused", params)
}
