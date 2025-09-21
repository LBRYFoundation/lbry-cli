package commands_transaction

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandTransactionShow() *cobra.Command {
	transaction_show := &cobra.Command{
		Use:   "show",
		Short: "Get a decoded transaction from a txid",
		Run:   HandleCommandTransactionShow,
	}

	transaction_show.Flags().String("txid", "", "(str) txid of the transaction")

	return transaction_show
}

func HandleCommandTransactionShow(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["txid"]
		if exists {
			cmd.Help()
			return
		}
		params["txid"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	_, exists := params["txid"]
	if !exists {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("transaction_show", params)
}
