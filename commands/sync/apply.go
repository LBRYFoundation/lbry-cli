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
	password := ""
	//password, _ := cmd.Flags().GetString("password")

	data, _ := cmd.Flags().GetString("data")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	blocking, _ := cmd.Flags().GetBool("blocking")

	// Check for arguments
	if len(args) >= 1 {
		password = args[0]
	}

	if password == "" {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{
		"password": password,
	}
	if data != "" {
		params["data"] = data
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if blocking {
		params["blocking"] = blocking
	}

	rpc.ExecuteRPCCommand("sync_apply", params)
}
