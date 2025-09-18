package commands

import (
	commands "lbry/cli/commands/sync"

	"github.com/spf13/cobra"
)

func CreateCommandSync() *cobra.Command {
	sync := &cobra.Command{
		Use:   "sync",
		Short: "Wallet synchronization.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	sync.AddCommand(commands.CreateCommandSyncApply())
	sync.AddCommand(commands.CreateCommandSyncHash())

	return sync
}
