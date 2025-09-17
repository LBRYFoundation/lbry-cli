package commands

import commands "lbry/cli/commands/collection"
import "github.com/spf13/cobra"

func CreateCommandCollection() *cobra.Command {
	collection := &cobra.Command{
		Use:   "collection",
		Short: "Create, update, list, resolve, and abandon collections.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	collection.AddCommand(commands.CreateCommandCollectionAbandon())
	collection.AddCommand(commands.CreateCommandCollectionCreate())
	collection.AddCommand(commands.CreateCommandCollectionList())
	collection.AddCommand(commands.CreateCommandCollectionResolve())
	collection.AddCommand(commands.CreateCommandCollectionUpdate())

	return collection
}
