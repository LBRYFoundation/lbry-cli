package commands_claim

import "github.com/spf13/cobra"

func CreateCommandClaimList() *cobra.Command {
	claim_list := &cobra.Command{
		Use:   "list",
		Short: "List my stream and channel claims.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return claim_list
}
