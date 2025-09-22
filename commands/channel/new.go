package commands_channel

import "github.com/spf13/cobra"

func CreateCommandChannelNew() *cobra.Command {
	channel_new := &cobra.Command{
		Use:        "new",
		Short:      "deprecated",
		Deprecated: "channel_new is deprecated, using channel_create.",
	}

	return channel_new
}
