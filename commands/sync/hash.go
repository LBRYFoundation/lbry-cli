package commands_sync

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSyncHash() *cobra.Command {
	sync_hash := &cobra.Command{
		Use:   "hash",
		Short: "Deterministic hash of the wallet.",
		Run:   HandleCommandSyncHash,
	}

	sync_hash.Flags().String("wallet_id", "", "(str) wallet for which to generate hash")

	return sync_hash
}

func HandleCommandSyncHash(cmd *cobra.Command, args []string) {
	wallet_id, _ := cmd.Flags().GetString("wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		wallet_id = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}

	rpc.ExecuteRPCCommand("sync_hash", params)
}
