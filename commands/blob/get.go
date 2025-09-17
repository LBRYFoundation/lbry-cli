package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobGet() *cobra.Command {
	blob_get := &cobra.Command{
		Use:   "get",
		Short: "Download and return a blob",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_get
}
