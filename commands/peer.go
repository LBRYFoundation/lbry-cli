package commands

import (
	commands "lbry/cli/commands/peer"

	"github.com/spf13/cobra"
)

func CreateCommandPeer() *cobra.Command {
	peer := &cobra.Command{
		Use:   "peer",
		Short: "DHT / Blob Exchange peer commands.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	peer.AddCommand(commands.CreateCommandPeerList())
	peer.AddCommand(commands.CreateCommandPeerPing())

	return peer
}
