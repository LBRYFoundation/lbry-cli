package commands_file

import "github.com/spf13/cobra"

func CreateCommandFileReflect() *cobra.Command {
	file_reflect := &cobra.Command{
		Use:   "reflect",
		Short: "Reflect all the blobs in a file matching the filter criteria",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file_reflect
}
