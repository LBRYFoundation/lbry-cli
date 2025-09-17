package commands_support

import "github.com/spf13/cobra"

func CreateCommandSupportList() *cobra.Command {
	support_list := &cobra.Command{
		Use:   "list",
		Short: "List staked supports and sent/received tips.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return support_list
}
