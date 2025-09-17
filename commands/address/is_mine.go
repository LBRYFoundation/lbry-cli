package commands_address

import "github.com/spf13/cobra"

func CreateCommandAddressIsMine() *cobra.Command {
	address_is_mine := &cobra.Command{
		Use:   "is_mine",
		Short: "Checks if an address is associated with the current wallet.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return address_is_mine
}
