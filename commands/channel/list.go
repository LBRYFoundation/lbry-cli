package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelList() *cobra.Command {
	channel_list := &cobra.Command{
		Use:   "list",
		Short: "List my channel claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_list
}
