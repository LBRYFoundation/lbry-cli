package commands

import (
	commands "lbry/cli/internal/commands/wallet"

	"github.com/spf13/cobra"
)

func CreateCommandWallet() *cobra.Command {
	wallet := &cobra.Command{
		Use:   "wallet",
		Short: "Create, modify and inspect wallets.",
	}

	wallet.AddCommand(commands.CreateCommandWalletAdd())
	wallet.AddCommand(commands.CreateCommandWalletBalance())
	wallet.AddCommand(commands.CreateCommandWalletCreate())
	wallet.AddCommand(commands.CreateCommandWalletDecrypt())
	wallet.AddCommand(commands.CreateCommandWalletEncrypt())
	wallet.AddCommand(commands.CreateCommandWalletExport())
	wallet.AddCommand(commands.CreateCommandWalletImport())
	wallet.AddCommand(commands.CreateCommandWalletList())
	wallet.AddCommand(commands.CreateCommandWalletLock())
	wallet.AddCommand(commands.CreateCommandWalletReconnect())
	wallet.AddCommand(commands.CreateCommandWalletRemove())
	wallet.AddCommand(commands.CreateCommandWalletSend())
	wallet.AddCommand(commands.CreateCommandWalletStatus())
	wallet.AddCommand(commands.CreateCommandWalletUnlock())

	return wallet
}
