package commands

import "github.com/spf13/cobra"

func CreateCommandAccount() *cobra.Command {
	account := &cobra.Command{
		Use:   "account",
		Short: "Create, modify and inspect wallet accounts.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	account.AddCommand(CreateCommandAccountAdd())
	account.AddCommand(CreateCommandAccountBalance())
	account.AddCommand(CreateCommandAccountCreate())
	account.AddCommand(CreateCommandAccountDeposit())
	account.AddCommand(CreateCommandAccountFund())
	account.AddCommand(CreateCommandAccountList())
	account.AddCommand(CreateCommandAccountMaxAddressGap())
	account.AddCommand(CreateCommandAccountRemove())
	account.AddCommand(CreateCommandAccountSend())
	account.AddCommand(CreateCommandAccountSet())

	return account
}
