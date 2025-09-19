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
	txid, _ := cmd.Flags().GetString("txid")

	// Check for arguments
	if len(args) >= 1 {
		txid = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	if txid == "" {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{}
	if txid != "" {
		params["txid"] = txid
	}

	rpc.ExecuteRPCCommand("transaction_show", params)
}
