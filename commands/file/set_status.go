package commands_file

import "github.com/spf13/cobra"

func CreateCommandFileSetStatus() *cobra.Command {
	file_set_status := &cobra.Command{
		Use:   "set_status",
		Short: "Start or stop downloading a file",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file_set_status
}
