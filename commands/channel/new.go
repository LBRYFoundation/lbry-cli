package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelNew() *cobra.Command {
	channel_new := &cobra.Command{
		Use:   "new",
		Short: "deprecated",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_new
}
