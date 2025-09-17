package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamUpdate() *cobra.Command {
	stream_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing stream claim and if a new file is provided announce it to lbrynet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_update
}
