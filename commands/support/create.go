package commands_support

import "github.com/spf13/cobra"

func CreateCommandSupportCreate() *cobra.Command {
	support_create := &cobra.Command{
		Use:   "create",
		Short: "Create a support or a tip for name claim.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return support_create
}
