package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelImport() *cobra.Command {
	channel_import := &cobra.Command{
		Use:   "import",
		Short: "Import serialized channel private key (to allow signing new streams to the channel)",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_import
}
