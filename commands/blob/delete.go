package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobDelete() *cobra.Command {
	blob_delete := &cobra.Command{
		Use:   "delete",
		Short: "Delete a blob",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_delete
}
