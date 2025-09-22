package commands_collection

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandCollectionAbandon() *cobra.Command {
	collection_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon one of my collection claims.",
		Run:   HandleCommandCollectionAbandon,
	}

	collection_abandon.Flags().String("claim_id", "", "(str) claim_id of the claim to abandon")
	collection_abandon.Flags().String("txid", "", "(str) txid of the claim to abandon")
	collection_abandon.Flags().Int("nout", -1, "(int) nout of the claim to abandon")
	collection_abandon.Flags().String("account_id", "", "(str) id of the account to use")
	collection_abandon.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	collection_abandon.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	collection_abandon.Flags().Bool("blocking", false, "(bool) wait until abandon is in mempool")

	return collection_abandon
}

func HandleCommandCollectionAbandon(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["claim_id"]
		if exists {
			cmd.Help()
			return
		}
		params["claim_id"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["txid"]
		if exists {
			cmd.Help()
			return
		}
		params["txid"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["nout"]
		if exists {
			cmd.Help()
			return
		}
		params["nout"] = args[2]
	}
	if len(args) > 3 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("collection_abandon", params)
}
