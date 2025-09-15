package commands

import "github.com/spf13/cobra"

func CreateCommandAccountBalance() *cobra.Command {
	account_balance := &cobra.Command{
		Use: "balance",
	}

	return account_balance
}
