package commands

import (
	commands "lbry/cli/internal/commands/peer"

	"github.com/spf13/cobra"
)

func CreateCommandPeer() *cobra.Command {
	peer := &cobra.Command{
		Use:   "peer",
		Short: "DHT / Blob Exchange peer commands.",
	}

	peer.AddCommand(commands.CreateCommandPeerList())
	peer.AddCommand(commands.CreateCommandPeerPing())

	return peer
}
