package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelCreate() *cobra.Command {
	channel_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new channel by generating a channel private key and establishing an '@' prefixed claim.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_create
}
