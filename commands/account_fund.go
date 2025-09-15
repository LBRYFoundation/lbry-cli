package commands

import "github.com/spf13/cobra"

func CreateCommandAccountFund() *cobra.Command {
	account_fund := &cobra.Command{
		Use: "fund",
	}

	return account_fund
}
