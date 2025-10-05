package commands_account

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountFund() *cobra.Command {
	account_fund := &cobra.Command{
		Use:   "fund",
		Short: "Transfer some amount (or --everything) to an account from another account (can be the same account). Amounts are interpreted as LBC. You can also spread the transfer across a number of --outputs (cannot be used together with --everything).",
		Run:   HandleCommandAccountFund,
	}

	account_fund.Flags().String("to_account", "", "(str) send to this account")
	account_fund.Flags().String("from_account", "", "(str) spend from this account")
	account_fund.Flags().Float64("amount", -1, "(decimal) the amount to transfer lbc")
	account_fund.Flags().Bool("everything", false, "(bool) transfer everything (excluding claims), default: false.")
	account_fund.Flags().Int("outputs", -1, "(int) split payment across many outputs, default: 1.")
	account_fund.Flags().String("wallet_id", "", "(str) limit operation to specific wallet.")
	account_fund.Flags().Bool("broadcast", false, "(bool) actually broadcast the transaction, default: false.")

	return account_fund
}

func HandleCommandAccountFund(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "to_account")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "from_account")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "everything")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "outputs")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "broadcast")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["to_account"]
		if exists {
			cmd.Help()
			return
		}
		params["to_account"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["from_account"]
		if exists {
			cmd.Help()
			return
		}
		params["from_account"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["amount"]
		if exists {
			cmd.Help()
			return
		}
		params["amount"] = args[2]
	}
	if len(args) >= 4 {
		_, exists := params["outputs"]
		if exists {
			cmd.Help()
			return
		}
		params["outputs"] = args[3]
	}
	if len(args) > 4 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_fund", params)
}
