package commands

import "github.com/spf13/cobra"

func CreateCommandAccountDeposit() *cobra.Command {
	account_deposit := &cobra.Command{
		Use: "deposit",
	}

	return account_deposit
}
