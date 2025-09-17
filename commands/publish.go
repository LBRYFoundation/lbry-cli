package commands

import "github.com/spf13/cobra"

func CreateCommandPublish() *cobra.Command {
	publish := &cobra.Command{
		Use:   "publish",
		Short: "Create or replace a stream claim at a given name (use 'stream create/update' for more control).",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return publish
}
