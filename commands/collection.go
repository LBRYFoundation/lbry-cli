package commands

import "github.com/spf13/cobra"

func CreateCommandCollection() *cobra.Command {
	collection := &cobra.Command{
		Use:   "collection",
		Short: "Create, update, list, resolve, and abandon collections.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection
}
