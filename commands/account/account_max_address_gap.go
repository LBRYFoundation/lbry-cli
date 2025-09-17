package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountMaxAddressGap() *cobra.Command {
	account_max_address_gap := &cobra.Command{
		Use:   "max_address_gap",
		Short: "Finds ranges of consecutive addresses that are unused and returns the length of the longest such range: for change and receiving address chains. This is useful to figure out ideal values to set for 'receiving_gap' and 'change_gap' account settings.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_max_address_gap
}
