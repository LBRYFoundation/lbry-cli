package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelUpdate() *cobra.Command {
	channel_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing channel claim.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_update
}
