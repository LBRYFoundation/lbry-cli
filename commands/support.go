package commands

import "github.com/spf13/cobra"

func CreateCommandSupport() *cobra.Command {
	support := &cobra.Command{
		Use:   "support",
		Short: "Create, list and abandon all types of supports.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return support
}
