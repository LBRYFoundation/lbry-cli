package commands_account

import "github.com/spf13/cobra"

func CreateCommandAccountSet() *cobra.Command {
	account_set := &cobra.Command{
		Use:   "set",
		Short: "Change various settings on an account.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_set
}
