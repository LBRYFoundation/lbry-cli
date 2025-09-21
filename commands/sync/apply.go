package commands_sync

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSyncApply() *cobra.Command {
	sync_apply := &cobra.Command{
		Use:   "apply",
		Short: "Apply incoming synchronization data, if provided, and return a sync hash and update wallet data.",
		Run:   HandleCommandSyncApply,
	}

	sync_apply.Flags().String("password", "", "(str) password to decrypt incoming and encrypt outgoing data")
	sync_apply.Flags().String("data", "", "(str) incoming sync data, if any")
	sync_apply.Flags().String("wallet_id", "", "(str) wallet being sync'ed")
	sync_apply.Flags().Bool("blocking", false, "(bool) wait until any new accounts have sync'ed")

	return sync_apply
}

func HandleCommandSyncApply(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "password")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "data")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["password"]
		if exists {
			cmd.Help()
			return
		}
		params["password"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("sync_apply", params)
}
