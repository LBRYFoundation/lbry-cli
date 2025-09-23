package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountBalance() *cobra.Command {
	account_balance := &cobra.Command{
		Use:   "balance",
		Short: "Return the balance of an account",
		Run:   HandleCommandAccountBalance,
	}

	account_balance.Flags().String("account_id", "", "(str) If provided only the balance for this account will be given. Otherwise default account.")
	account_balance.Flags().String("address", "", "")
	account_balance.Flags().String("wallet_id", "", "(str) balance for specific wallet")
	account_balance.Flags().Int("confirmations", -1, "(int) Only include transactions with this many confirmed blocks.")

	return account_balance
}

func HandleCommandAccountBalance(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "confirmations")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["account_id"]
		if exists {
			cmd.Help()
			return
		}
		params["account_id"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["address"]
		if exists {
			cmd.Help()
			return
		}
		params["address"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["confirmations"]
		if exists {
			cmd.Help()
			return
		}
		params["confirmations"] = args[2]
	}
	if len(args) > 3 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("account_balance", params)
}
