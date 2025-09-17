package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobReflectAll() *cobra.Command {
	blob_reflect_all := &cobra.Command{
		Use:   "reflect_all",
		Short: "Reflects all saved blobs",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_reflect_all
}
