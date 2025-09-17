package commands

import "github.com/spf13/cobra"

func CreateCommandGet() *cobra.Command {
	get := &cobra.Command{
		Use:   "get",
		Short: "Download stream from a LBRY name.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return get
}
