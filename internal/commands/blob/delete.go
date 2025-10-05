package commands_blob

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobDelete() *cobra.Command {
	blob_delete := &cobra.Command{
		Use:   "delete",
		Short: "Delete a blob",
		Run:   HandleCommandBlobDelete,
	}

	blob_delete.Flags().String("blob_hash", "", "(str) blob hash of the blob to delete")

	return blob_delete
}

func HandleCommandBlobDelete(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "blob_hash")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["blob_hash"]
		if exists {
			cmd.Help()
			return
		}
		params["blob_hash"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "blob_delete", params)
}
