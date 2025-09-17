package commands

import commands "lbry/cli/commands/address"
import "github.com/spf13/cobra"

func CreateCommandAddress() *cobra.Command {
	address := &cobra.Command{
		Use:   "address",
		Short: "List, generate and verify addresses.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	address.AddCommand(commands.CreateCommandAddressIsMine())
	address.AddCommand(commands.CreateCommandAddressList())
	address.AddCommand(commands.CreateCommandAddressUnused())

	return address
}
