package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamList() *cobra.Command {
	stream_list := &cobra.Command{
		Use:   "list",
		Short: "List my stream claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_list
}
