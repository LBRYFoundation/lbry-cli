package commands

import (
	commands "lbry/cli/internal/commands/claim"

	"github.com/spf13/cobra"
)

func CreateCommandClaim() *cobra.Command {
	claim := &cobra.Command{
		Use:   "claim",
		Short: "List and search all types of claims.",
	}

	claim.AddCommand(commands.CreateCommandClaimSearch())

	return claim
}
