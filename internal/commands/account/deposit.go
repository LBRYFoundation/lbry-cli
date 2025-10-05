package commands_account

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountDeposit() *cobra.Command {
	account_deposit := &cobra.Command{
		Use:   "deposit",
		Short: "Spend a time locked transaction into your account.",
		Run:   HandleCommandAccountDeposit,
	}

	account_deposit.Flags().String("txid", "", "(str) id of the transaction")
	account_deposit.Flags().Int("nout", -1, "(int) output number in the transaction")
	account_deposit.Flags().String("redeem_script", "", "(str) redeem script for output")
	account_deposit.Flags().String("private_key", "", "(str) private key to sign transaction")
	account_deposit.Flags().String("to_account", "", "(str) deposit to this account")
	account_deposit.Flags().String("wallet_id", "", "(str) limit operation to specific wallet.")
	account_deposit.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	account_deposit.Flags().Bool("blocking", false, "(bool) wait until tx has synced")

	return account_deposit
}

func HandleCommandAccountDeposit(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "redeem_script")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "private_key")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "to_account")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["txid"]
		if exists {
			cmd.Help()
			return
		}
		params["txid"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["nout"]
		if exists {
			cmd.Help()
			return
		}
		params["nout"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["redeem_script"]
		if exists {
			cmd.Help()
			return
		}
		params["redeem_script"] = args[2]
	}
	if len(args) >= 4 {
		_, exists := params["private_key"]
		if exists {
			cmd.Help()
			return
		}
		params["private_key"] = args[3]
	}
	if len(args) >= 5 {
		_, exists := params["to_account"]
		if exists {
			cmd.Help()
			return
		}
		params["to_account"] = args[4]
	}
	if len(args) > 5 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, existsTXID := params["txid"]
	if !existsTXID {
		cmd.Help()
		return
	}
	_, existsNOut := params["nout"]
	if !existsNOut {
		cmd.Help()
		return
	}
	_, existsRedeemScript := params["redeem_script"]
	if !existsRedeemScript {
		cmd.Help()
		return
	}
	_, existsPrivateKey := params["private_key"]
	if !existsPrivateKey {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_deposit", params)
}
