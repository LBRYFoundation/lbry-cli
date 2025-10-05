package commands_blob

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobGet() *cobra.Command {
	blob_get := &cobra.Command{
		Use:   "get",
		Short: "Download and return a blob",
		Run:   HandleCommandBlobGet,
	}

	blob_get.Flags().String("blob_hash", "", "(str) blob hash of the blob to get")
	blob_get.Flags().Int("timeout", -1, "(int) timeout in number of seconds")
	blob_get.Flags().Bool("read", false, "")

	return blob_get
}

func HandleCommandBlobGet(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "blob_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "timeout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "read")

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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "blob_get", params)
}
