package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountList() *cobra.Command {
	account_list := &cobra.Command{
		Use:   "list",
		Short: "List details of all of the accounts or a specific account.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_list
}
