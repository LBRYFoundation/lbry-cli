package commands

import (
	commands "lbry/cli/commands/sync"

	"github.com/spf13/cobra"
)

func CreateCommandSync() *cobra.Command {
	sync := &cobra.Command{
		Use:   "sync",
		Short: "Wallet synchronization.",
	}

	sync.AddCommand(commands.CreateCommandSyncApply())
	sync.AddCommand(commands.CreateCommandSyncHash())

	return sync
}
