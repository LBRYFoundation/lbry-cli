package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountRemove() *cobra.Command {
	account_remove := &cobra.Command{
		Use:   "remove",
		Short: "Remove an existing account.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return account_remove
}
