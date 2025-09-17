package commands_sync

import "github.com/spf13/cobra"

func CreateCommandSyncApply() *cobra.Command {
	sync_apply := &cobra.Command{
		Use:   "apply",
		Short: "Apply incoming synchronization data, if provided, and return a sync hash and update wallet data.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return sync_apply
}
