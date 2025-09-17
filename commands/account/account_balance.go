package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountBalance() *cobra.Command {
	account_balance := &cobra.Command{
		Use:   "balance",
		Short: "Return the balance of an account",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_balance
}
