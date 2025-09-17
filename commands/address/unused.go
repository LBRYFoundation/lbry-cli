package commands_address

import "github.com/spf13/cobra"

func CreateCommandAddressUnused() *cobra.Command {
	address_unused := &cobra.Command{
		Use:   "unused",
		Short: "Return an address containing no balance, will create a new address if there is none.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return address_unused
}
