package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamRepost() *cobra.Command {
	stream_repost := &cobra.Command{
		Use:   "repost",
		Short: "Creates a claim that references an existing stream by its claim id.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_repost
}
