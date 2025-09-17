package commands_file

import "github.com/spf13/cobra"

func CreateCommandFileList() *cobra.Command {
	file_list := &cobra.Command{
		Use:   "list",
		Short: "List files limited by optional filters",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file_list
}
