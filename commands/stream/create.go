package commands_stream

import "github.com/spf13/cobra"

func CreateCommandStreamCreate() *cobra.Command {
	stream_create := &cobra.Command{
		Use:   "create",
		Short: "Make a new stream claim and announce the associated file to lbrynet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream_create
}
