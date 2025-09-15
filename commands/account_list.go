package commands

import "github.com/spf13/cobra"

func CreateCommandAccountList() *cobra.Command {
	account_list := &cobra.Command{
		Use: "list",
	}

	return account_list
}
