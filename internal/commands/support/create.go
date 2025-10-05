package commands_support

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSupportCreate() *cobra.Command {
	support_create := &cobra.Command{
		Use:   "create",
		Short: "Create a support or a tip for name claim.",
		Run:   HandleCommandSupportCreate,
	}

	support_create.Flags().String("claim_id", "", "(str) claim_id of the claim to support")
	support_create.Flags().Float64("amount", -1, "(decimal) amount of support")
	support_create.Flags().Bool("tip", false, "(bool) send support to claim owner, default: false.")
	support_create.Flags().String("channel_id", "", "(str) claim id of the supporters identity channel")
	support_create.Flags().String("channel_name", "", "(str) name of the supporters identity channel")
	support_create.Flags().String("channel_account_id", "", "(str) one or more account ids for accounts to look in for channel certificates, defaults to all accounts.")
	support_create.Flags().String("account_id", "", "(str) account to use for holding the transaction")
	support_create.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	support_create.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	support_create.Flags().String("comment", "", "(str) add a comment to the support")
	support_create.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	support_create.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")

	return support_create
}

func HandleCommandSupportCreate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "tip")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "comment")
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
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "support_create", params)
}
