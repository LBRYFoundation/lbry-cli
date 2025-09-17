package commands_collection

import "github.com/spf13/cobra"

func CreateCommandCollectionCreate() *cobra.Command {
	collection_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new collection.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection_create
}
