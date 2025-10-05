package commands_support

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSupportAbandon() *cobra.Command {
	support_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon supports, including tips, of a specific claim, optionally keeping some amount as supports.",
		Run:   HandleCommandSupportAbandon,
	}

	support_abandon.Flags().String("claim_id", "", "(str) claim_id of the support to abandon")
	support_abandon.Flags().String("txid", "", "(str) txid of the claim to abandon")
	support_abandon.Flags().Int("nout", -1, "(int) nout of the claim to abandon")
	support_abandon.Flags().Float64("keep", -1, "(decimal) amount of lbc to keep as support")
	support_abandon.Flags().String("account_id", "", "(str) id of the account to use")
	support_abandon.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	support_abandon.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	support_abandon.Flags().Bool("blocking", false, "(bool) wait until abandon is in mempool")

	return support_abandon
}

func HandleCommandSupportAbandon(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "keep")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "support_abandon", params)
}
