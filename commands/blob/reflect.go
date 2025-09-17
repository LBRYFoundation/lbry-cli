package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobReflect() *cobra.Command {
	blob_reflect := &cobra.Command{
		Use:   "reflect",
		Short: "Reflects specified blobs",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_reflect
}
