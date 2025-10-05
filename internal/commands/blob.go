package commands

import (
	commands "lbry/cli/internal/commands/blob"

	"github.com/spf13/cobra"
)

func CreateCommandBlob() *cobra.Command {
	blob := &cobra.Command{
		Use:   "blob",
		Short: "Blob management.",
	}

	blob.AddCommand(commands.CreateCommandBlobAnnounce())
	blob.AddCommand(commands.CreateCommandBlobClean())
	blob.AddCommand(commands.CreateCommandBlobDelete())
	blob.AddCommand(commands.CreateCommandBlobGet())
	blob.AddCommand(commands.CreateCommandBlobList())
	blob.AddCommand(commands.CreateCommandBlobReflect())
	blob.AddCommand(commands.CreateCommandBlobReflectAll())

	return blob
}
