package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountFund() *cobra.Command {
	account_fund := &cobra.Command{
		Use:   "fund",
		Short: "Transfer some amount (or --everything) to an account from another account (can be the same account). Amounts are interpreted as LBC. You can also spread the transfer across a number of --outputs (cannot be used together with --everything).",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_fund
}
