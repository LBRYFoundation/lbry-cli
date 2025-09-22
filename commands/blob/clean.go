package commands_blob

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobClean() *cobra.Command {
	blob_clean := &cobra.Command{
		Use:   "clean",
		Short: "Deletes blobs to cleanup disk space",
		Run:   HandleCommandBlobClean,
	}

	return blob_clean
}

func HandleCommandBlobClean(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("blob_clean")
}
