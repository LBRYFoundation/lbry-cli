package commands_purchase

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPurchaseCreate() *cobra.Command {
	purchase_create := &cobra.Command{
		Use:   "create",
		Short: "Purchase a claim.",
		Run:   HandleCommandPurchaseCreate,
	}

	purchase_create.Flags().String("claim_id", "", "(str) claim id of claim to purchase")
	purchase_create.Flags().String("url", "", "(str) lookup claim to purchase by url")
	purchase_create.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	purchase_create.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	purchase_create.Flags().Bool("allow_duplicate_purchase", false, "(bool) allow purchasing claim_id you already own")
	purchase_create.Flags().Bool("override_max_key_fee", false, "(bool) ignore max key fee for this purchase")
	purchase_create.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	purchase_create.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")

	return purchase_create
}

func HandleCommandPurchaseCreate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "allow_duplicate_purchase")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "override_max_key_fee")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "purchase_create", params)
}
