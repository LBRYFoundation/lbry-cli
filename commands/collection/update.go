package commands_collection

import "github.com/spf13/cobra"

func CreateCommandCollectionUpdate() *cobra.Command {
	collection_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing collection claim.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection_update
}
