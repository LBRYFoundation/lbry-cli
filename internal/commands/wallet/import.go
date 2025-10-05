package commands_wallet

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletImport() *cobra.Command {
	wallet_import := &cobra.Command{
		Use:   "import",
		Short: "Import wallet data and merge accounts and preferences. Data is expected to be JSON if password is not supplied.",
		Run:   HandleCommandWalletImport,
	}

	wallet_import.Flags().String("data", "", "(str) incoming wallet data")
	wallet_import.Flags().String("password", "", "(str) password to decrypt incoming data")
	wallet_import.Flags().String("wallet_id", "", "(str) wallet being merged into")
	wallet_import.Flags().Bool("blocking", false, "(bool) wait until any new accounts have merged")

	return wallet_import
}

func HandleCommandWalletImport(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "data")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "password")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["data"]
		if exists {
			cmd.Help()
			return
		}
		params["data"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["password"]
		if exists {
			cmd.Help()
			return
		}
		params["password"] = args[0]
	}
	if len(args) > 2 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_import", params)
}
