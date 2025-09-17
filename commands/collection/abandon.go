package commands_collection

import "github.com/spf13/cobra"

func CreateCommandCollectionAbandon() *cobra.Command {
	collection_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon one of my collection claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return collection_abandon
}
