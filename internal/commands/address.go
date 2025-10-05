package commands

import (
	commands "lbry/cli/internal/commands/address"

	"github.com/spf13/cobra"
)

func CreateCommandAddress() *cobra.Command {
	address := &cobra.Command{
		Use:   "address",
		Short: "List, generate and verify addresses.",
	}

	address.AddCommand(commands.CreateCommandAddressIsMine())
	address.AddCommand(commands.CreateCommandAddressList())
	address.AddCommand(commands.CreateCommandAddressUnused())

	return address
}
