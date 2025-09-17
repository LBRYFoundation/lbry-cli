package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountDeposit() *cobra.Command {
	account_deposit := &cobra.Command{
		Use:   "deposit",
		Short: "Spend a time locked transaction into your account.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_deposit
}
