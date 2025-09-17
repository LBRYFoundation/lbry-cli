package commands_sync

import "github.com/spf13/cobra"

func CreateCommandSyncHash() *cobra.Command {
	sync_hash := &cobra.Command{
		Use:   "hash",
		Short: "Deterministic hash of the wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return sync_hash
}
