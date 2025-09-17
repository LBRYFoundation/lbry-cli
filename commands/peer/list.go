package commands_peer

import "github.com/spf13/cobra"

func CreateCommandPeerList() *cobra.Command {
	peer_list := &cobra.Command{
		Use:   "list",
		Short: "Get peers for blob hash",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return peer_list
}
