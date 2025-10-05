package commands_blob

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobReflectAll() *cobra.Command {
	blob_reflect_all := &cobra.Command{
		Use:   "reflect_all",
		Short: "Reflects all saved blobs",
		Run:   HandleCommandBlobReflectAll,
	}

	return blob_reflect_all
}

func HandleCommandBlobReflectAll(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "blob_reflect_all")
}
