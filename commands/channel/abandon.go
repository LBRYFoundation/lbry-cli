package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelAbandon() *cobra.Command {
	channel_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon one of my channel claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_abandon
}
