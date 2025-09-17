package commands_claim

import "github.com/spf13/cobra"

func CreateCommandClaimSearch() *cobra.Command {
	claim_search := &cobra.Command{
		Use:   "search",
		Short: "Search for stream and channel claims on the blockchain.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return claim_search
}
