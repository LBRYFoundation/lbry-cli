package commands_account

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountSend() *cobra.Command {
	account_send := &cobra.Command{
		Use:   "send",
		Short: "Send the same number of credits to multiple addresses from a specific account (or default account).",
		Run:   HandleCommandAccountSend,
	}

	//account_send.Flags().String("amount", "", "")
	//account_send.Flags().String("addresses", nil, "")
	account_send.Flags().String("account_id", "", "(str) account to fund the transaction")
	account_send.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	account_send.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	account_send.Flags().Bool("blocking", false, "(bool) wait until tx has synced")

	return account_send
}

func HandleCommandAccountSend(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "amount")
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "addresses")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["amount"]
		if exists {
			cmd.Help()
			return
		}
		params["amount"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["addresses"]
		if exists {
			cmd.Help()
			return
		}
		params["addresses"] = args[1:]
	}

	// Check for required parameters
	_, existsAmount := params["amount"]
	if !existsAmount {
		cmd.Help()
		return
	}
	_, existsAddresses := params["addresses"]
	if !existsAddresses {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("account_send", params)
}
