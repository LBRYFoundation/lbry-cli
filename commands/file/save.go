package commands_file

import "github.com/spf13/cobra"

func CreateCommandFileSave() *cobra.Command {
	file_save := &cobra.Command{
		Use:   "save",
		Short: "Start saving a file to disk.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return file_save
}
