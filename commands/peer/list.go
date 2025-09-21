package commands_peer

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPeerList() *cobra.Command {
	peer_list := &cobra.Command{
		Use:   "list",
		Short: "Get peers for blob hash",
		Run:   HandleCommandPeerList,
	}

	peer_list.Flags().String("blob_hash", "", "(str) find available peers for this blob hash")
	peer_list.Flags().Int("page", -1, "(int) page to return during paginating")
	peer_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return peer_list
}

func HandleCommandPeerList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "blob_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

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

	rpc.ExecuteRPCCommand("peer_list", params)
}
