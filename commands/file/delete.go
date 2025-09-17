package commands_file

import "github.com/spf13/cobra"

func CreateCommandFileDelete() *cobra.Command {
	file_delete := &cobra.Command{
		Use:   "file",
		Short: "Delete a LBRY file",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file_delete
}
