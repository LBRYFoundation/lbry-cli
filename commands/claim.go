package commands

import commands "lbry/cli/commands/claim"
import "github.com/spf13/cobra"

func CreateCommandClaim() *cobra.Command {
	claim := &cobra.Command{
		Use:   "claim",
		Short: "List and search all types of claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	claim.AddCommand(commands.CreateCommandClaimSearch())

	return claim
}
