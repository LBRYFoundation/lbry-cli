package commands

import "github.com/spf13/cobra"

func CreateCommandAccountSend() *cobra.Command {
	account_send := &cobra.Command{
		Use: "send",
	}

	return account_send
}
