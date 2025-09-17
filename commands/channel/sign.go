package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelSign() *cobra.Command {
	channel_sign := &cobra.Command{
		Use:   "sign",
		Short: "Signs data using the specified channel signing key.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_sign
}
