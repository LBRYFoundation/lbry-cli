package commands

import (
	commands "lbry/cli/internal/commands/account"

	"github.com/spf13/cobra"
)

func CreateCommandAccount() *cobra.Command {
	account := &cobra.Command{
		Use:   "account",
		Short: "Create, modify and inspect wallet accounts.",
	}

	account.AddCommand(commands.CreateCommandAccountAdd())
	account.AddCommand(commands.CreateCommandAccountBalance())
	account.AddCommand(commands.CreateCommandAccountCreate())
	account.AddCommand(commands.CreateCommandAccountDeposit())
	account.AddCommand(commands.CreateCommandAccountFund())
	account.AddCommand(commands.CreateCommandAccountList())
	account.AddCommand(commands.CreateCommandAccountMaxAddressGap())
	account.AddCommand(commands.CreateCommandAccountRemove())
	account.AddCommand(commands.CreateCommandAccountSend())
	account.AddCommand(commands.CreateCommandAccountSet())

	return account
}
