package commands

import "github.com/spf13/cobra"

func CreateCommandResolve() *cobra.Command {
	resolve := &cobra.Command{
		Use:   "resolve",
		Short: "Get the claim that a URL refers to.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return resolve
}
