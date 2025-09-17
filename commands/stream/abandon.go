package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamAbandon() *cobra.Command {
	stream_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon one of my stream claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_abandon
}
