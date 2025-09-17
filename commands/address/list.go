package commands_address

import "github.com/spf13/cobra"

func CreateCommandAddressList() *cobra.Command {
	address_list := &cobra.Command{
		Use:   "list",
		Short: "List account addresses or details of single address.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return address_list
}
