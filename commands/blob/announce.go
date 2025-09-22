package commands_blob

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobAnnounce() *cobra.Command {
	blob_announce := &cobra.Command{
		Use:   "announce",
		Short: "Announce blobs to the DHT",
		Run:   HandleCommandBlobAnnounce,
	}

	blob_announce.Flags().String("blob_hash", "", "(str) announce a blob, specified by blob_hash")
	blob_announce.Flags().String("stream_hash", "", "(str) announce all blobs associated with stream_hash")
	blob_announce.Flags().String("sd_hash", "", "(str) announce all blobs associated with sd_hash and the sd_hash itself")

	return blob_announce
}

func HandleCommandBlobAnnounce(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "blob_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")

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

	rpc.ExecuteRPCCommand("blob_announce", params)
}
