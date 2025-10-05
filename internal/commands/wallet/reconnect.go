package commands_wallet

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletReconnect() *cobra.Command {
	wallet_reconnect := &cobra.Command{
		Use:   "reconnect",
		Short: "Reconnects ledger network client, applying new configurations.",
		Run:   HandleCommandWalletReconnect,
	}

	return wallet_reconnect
}

func HandleCommandWalletReconnect(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_reconnect")
}
