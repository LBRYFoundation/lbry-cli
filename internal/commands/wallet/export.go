package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletExport() *cobra.Command {
	wallet_export := &cobra.Command{
		Use:   "export",
		Short: "Exports encrypted wallet data if password is supplied; otherwise plain JSON.",
		Run:   HandleCommandWalletExport,
	}

	wallet_export.Flags().String("password", "", "(str) password to encrypt outgoing data")
	wallet_export.Flags().String("wallet_id", "", "(str) wallet being exported")

	return wallet_export
}

func HandleCommandWalletExport(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "password")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_export", params)
}
