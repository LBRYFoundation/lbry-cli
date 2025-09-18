package commands

import (
	commands "lbry/cli/commands/support"

	"github.com/spf13/cobra"
)

func CreateCommandSupport() *cobra.Command {
	support := &cobra.Command{
		Use:   "support",
		Short: "Create, list and abandon all types of supports.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	support.AddCommand(commands.CreateCommandSupportAbandon())
	support.AddCommand(commands.CreateCommandSupportCreate())
	support.AddCommand(commands.CreateCommandSupportList())
	support.AddCommand(commands.CreateCommandSupportSum())

	return support
}
