package commands

import "github.com/spf13/cobra"

func CreateCommandStream() *cobra.Command {
	stream := &cobra.Command{
		Use:   "stream",
		Short: "Create, update, abandon, list and inspect your stream claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return stream
}
