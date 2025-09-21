package commands

import (
	commands "lbry/cli/commands/channel"

	"github.com/spf13/cobra"
)

func CreateCommandChannel() *cobra.Command {
	channel := &cobra.Command{
		Use:   "channel",
		Short: "Create, update, abandon and list your channel claims.",
	}

	channel.AddCommand(commands.CreateCommandChannelAbandon())
	channel.AddCommand(commands.CreateCommandChannelCreate())
	channel.AddCommand(commands.CreateCommandChannelExport())
	channel.AddCommand(commands.CreateCommandChannelImport())
	channel.AddCommand(commands.CreateCommandChannelList())
	channel.AddCommand(commands.CreateCommandChannelNew())
	channel.AddCommand(commands.CreateCommandChannelSign())
	channel.AddCommand(commands.CreateCommandChannelUpdate())

	return channel
}
