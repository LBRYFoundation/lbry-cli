package commands

import "github.com/spf13/cobra"

func CreateCommandAccountSend() *cobra.Command {
	account_send := &cobra.Command{
		Use:   "send",
		Short: "Send the same number of credits to multiple addresses from a specific account (or default account).",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	return account_send
}
