package commands_support

import "github.com/spf13/cobra"

func CreateCommandSupportAbandon() *cobra.Command {
	support_abandon := &cobra.Command{
		Use:   "abandon",
		Short: "Abandon supports, including tips, of a specific claim, optionally keeping some amount as supports.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return support_abandon
}
