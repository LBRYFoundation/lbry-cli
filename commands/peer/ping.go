package commands_peer

import "github.com/spf13/cobra"

func CreateCommandPeerPing() *cobra.Command {
	peer_ping := &cobra.Command{
		Use:   "ping",
		Short: "Send a kademlia ping to the specified peer. If address and port are provided the peer is directly pinged, if not provided the peer is located first.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return peer_ping
}
