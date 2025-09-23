package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletCreate() *cobra.Command {
	wallet_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new wallet.",
		Run:   HandleCommandWalletCreate,
	}

	wallet_create.Flags().String("wallet_id", "", "(str) wallet file name")
	wallet_create.Flags().Bool("skip_on_startup", false, "(bool) don't add wallet to daemon_settings.yml")
	wallet_create.Flags().Bool("create_account", false, "(bool) generates the default account")
	wallet_create.Flags().Bool("single_key", false, "(bool) used with --create_account, creates single-key account")

	return wallet_create
}

func HandleCommandWalletCreate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "skip_on_startup")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "create_account")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "single_key")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["wallet_id"]
		if exists {
			cmd.Help()
			return
		}
		params["wallet_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("wallet_create", params)
}
