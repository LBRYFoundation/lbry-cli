package commands_blob

import "github.com/spf13/cobra"

func CreateCommandBlobAnnounce() *cobra.Command {
	blob_announce := &cobra.Command{
		Use:   "announce",
		Short: "Announce blobs to the DHT",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob_announce
}
