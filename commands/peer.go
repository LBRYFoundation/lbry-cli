package commands

import "github.com/spf13/cobra"

func CreateCommandPeer() *cobra.Command {
	peer := &cobra.Command{
		Use:   "peer",
		Short: "DHT / Blob Exchange peer commands.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return peer
}
