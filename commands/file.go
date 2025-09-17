package commands

import "github.com/spf13/cobra"

func CreateCommandFile() *cobra.Command {
	file := &cobra.Command{
		Use:   "file",
		Short: "File management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file
}
