package commands_blob

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobList() *cobra.Command {
	blob_list := &cobra.Command{
		Use:   "list",
		Short: "Returns blob hashes. If not given filters, returns all blobs known by the blob manager",
		Run:   HandleCommandBlobList,
	}

	blob_list.Flags().Bool("needed", false, "(bool) only return needed blobs")
	blob_list.Flags().Bool("finished", false, "(bool) only return finished blobs")
	blob_list.Flags().String("uri", "", "(str) filter blobs by stream in a uri")
	blob_list.Flags().String("stream_hash", "", "(str) filter blobs in a stream by sd hash, ie the hash of the stream descriptor blob for a stream that has been downloaded")
	blob_list.Flags().Int("page", -1, "(int) page to return during paginating")
	blob_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return blob_list
}

func HandleCommandBlobList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "needed")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "finished")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "uri")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["uri"]
		if exists {
			cmd.Help()
			return
		}
		params["uri"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["stream_hash"]
		if exists {
			cmd.Help()
			return
		}
		params["stream_hash"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["sd_hash"]
		if exists {
			cmd.Help()
			return
		}
		params["sd_hash"] = args[2]
	}
	if len(args) > 3 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "blob_list", params)
}
