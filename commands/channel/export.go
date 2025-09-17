package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelExport() *cobra.Command {
	channel_export := &cobra.Command{
		Use:   "export",
		Short: "Export channel private key.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return channel_export
}
