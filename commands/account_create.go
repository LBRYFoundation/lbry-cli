package commands

import "github.com/spf13/cobra"

func CreateCommandAccountCreate() *cobra.Command {
	account_create := &cobra.Command{
		Use: "create",
	}

	return account_create
}
