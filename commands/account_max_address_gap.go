package commands

import "github.com/spf13/cobra"

func CreateCommandAccountMaxAddressGap() *cobra.Command {
	account_max_address_gap := &cobra.Command{
		Use: "max_address_gap",
	}

	return account_max_address_gap
}
