package commands

import (
	commands "lbry/cli/commands/file"

	"github.com/spf13/cobra"
)

func CreateCommandFile() *cobra.Command {
	file := &cobra.Command{
		Use:   "file",
		Short: "File management.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	file.AddCommand(commands.CreateCommandFileDelete())
	file.AddCommand(commands.CreateCommandFileList())
	file.AddCommand(commands.CreateCommandFileReflect())
	file.AddCommand(commands.CreateCommandFileSave())
	file.AddCommand(commands.CreateCommandFileSetStatus())

	return file
}
