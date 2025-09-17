package commands

import "github.com/spf13/cobra"

func CreateCommandBlob() *cobra.Command {
	blob := &cobra.Command{
		Use:   "blob",
		Short: "Blob management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return blob
}
