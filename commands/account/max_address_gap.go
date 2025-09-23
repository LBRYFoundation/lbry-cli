package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountMaxAddressGap() *cobra.Command {
	account_max_address_gap := &cobra.Command{
		Use:   "max_address_gap",
		Short: "Finds ranges of consecutive addresses that are unused and returns the length of the longest such range: for change and receiving address chains. This is useful to figure out ideal values to set for 'receiving_gap' and 'change_gap' account settings.",
		Run:   HandleCommandAccountMaxAddressGap,
	}

	account_max_address_gap.Flags().String("account_id", "", "(str) account for which to get max gaps")
	account_max_address_gap.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return account_max_address_gap
}

func HandleCommandAccountMaxAddressGap(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

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

	rpc.ExecuteRPCCommand("account_max_address_gap", params)
}
