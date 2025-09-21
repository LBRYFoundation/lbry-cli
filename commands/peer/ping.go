package commands_peer

import (
	"lbry/cli/rpc"
	"strconv"

	"github.com/spf13/cobra"
)

func CreateCommandPeerPing() *cobra.Command {
	peer_ping := &cobra.Command{
		Use:   "ping",
		Short: "Send a kademlia ping to the specified peer. If address and port are provided the peer is directly pinged, if not provided the peer is located first.",
		Run:   HandleCommandPeerPing,
	}

	peer_ping.Flags().String("node_id", "", "")
	peer_ping.Flags().String("address", "", "")
	peer_ping.Flags().Int("port", -1, "")

	return peer_ping
}

func HandleCommandPeerPing(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "node_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "port")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["node_id"]
		if exists {
			cmd.Help()
			return
		}
		params["node_id"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["address"]
		if exists {
			cmd.Help()
			return
		}
		params["address"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["port"]
		if exists {
			cmd.Help()
			return
		}
		val, _ := strconv.Atoi(args[2])
		params["port"] = val
	}
	if len(args) > 3 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("peer_ping", params)
}
