package commands

import "github.com/spf13/cobra"

func CreateCommandChannel() *cobra.Command {
	channel := &cobra.Command{
		Use:   "channel",
		Short: "Create, update, abandon and list your channel claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel
}
