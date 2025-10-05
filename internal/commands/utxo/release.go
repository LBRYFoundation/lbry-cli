package commands_utxo

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandUTXORelease() *cobra.Command {
	utxo_release := &cobra.Command{
		Use:   "release",
		Short: "When spending a UTXO it is locally locked to prevent double spends; occasionally this can result in a UTXO being locked which ultimately did not get spent (failed to broadcast, spend transaction was not accepted by blockchain node, etc). This command releases the lock on all UTXOs in your account.",
		Run:   HandleCommandUTXORelease,
	}

	utxo_release.Flags().String("account_id", "", "(str) id of the account to query")
	utxo_release.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	utxo_release.Flags().Int("page", -1, "(int) page to return during paginating")
	utxo_release.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return utxo_release
}

func HandleCommandUTXORelease(cmd *cobra.Command, args []string) {
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "utxo_release", params)
}
