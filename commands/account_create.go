package commands

import "github.com/spf13/cobra"

func CreateCommandAccountCreate() *cobra.Command {
	account_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new account. Specify --single_key if you want to use the same address for all transactions (not recommended).",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_create
}
