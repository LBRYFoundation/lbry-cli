package commands_collection

import "github.com/spf13/cobra"

func CreateCommandCollectionResolve() *cobra.Command {
	collection_resolve := &cobra.Command{
		Use:   "resolve",
		Short: "Resolve claims in the collection.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection_resolve
}
