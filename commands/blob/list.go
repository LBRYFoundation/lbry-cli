package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobList() *cobra.Command {
	blob_list := &cobra.Command{
		Use:   "list",
		Short: "Returns blob hashes. If not given filters, returns all blobs known by the blob manager",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_list
}
