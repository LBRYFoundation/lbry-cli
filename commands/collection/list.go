package commands_collection

import "github.com/spf13/cobra"

func CreateCommandCollectionList() *cobra.Command {
	collection_list := &cobra.Command{
		Use:   "list",
		Short: "List my collection claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection_list
}
