package commands

import "github.com/spf13/cobra"

func CreateCommandAccountRemove() *cobra.Command {
	account_remove := &cobra.Command{
		Use: "remove",
	}

	return account_remove
}
