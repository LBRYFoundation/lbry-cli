package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobClean() *cobra.Command {
	blob_clean := &cobra.Command{
		Use:   "clean",
		Short: "Deletes blobs to cleanup disk space",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_clean
}
