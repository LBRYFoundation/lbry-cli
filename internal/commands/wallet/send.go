package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletSend() *cobra.Command {
	wallet_send := &cobra.Command{
		Use:   "send",
		Short: "Send the same number of credits to multiple addresses using all accounts in wallet to fund the transaction and the default account to receive any change.",
		Run:   HandleCommandWalletSend,
	}

	//wallet_send.Flags().String("amount", "", "")
	//wallet_send.Flags().StringArray("addresses", nil, "")
	wallet_send.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	wallet_send.Flags().String("change_account_id", "", "(str) account where change will go")
	wallet_send.Flags().String("funding_account_ids", "", "(str) accounts to fund the transaction")
	wallet_send.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	wallet_send.Flags().Bool("blocking", false, "(bool) wait until tx has synced")

	return wallet_send
}

func HandleCommandWalletSend(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "amount")
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "addresses")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "change_account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "funding_account_ids")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_send", params)
}
