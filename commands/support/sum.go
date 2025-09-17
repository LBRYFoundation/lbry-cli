package commands_support

import "github.com/spf13/cobra"

func CreateCommandSupportSum() *cobra.Command {
	support_sum := &cobra.Command{
		Use:   "sum",
		Short: "List total staked supports for a claim, grouped by the channel that signed the support.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return support_sum
}
