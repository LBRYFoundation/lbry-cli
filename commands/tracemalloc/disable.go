package commands_tracemalloc

import "github.com/spf13/cobra"

func CreateCommandTraceMAllocDisable() *cobra.Command {
	tracemalloc_disable := &cobra.Command{
		Use:   "disable",
		Short: "Disable tracemalloc memory tracing",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return tracemalloc_disable
}
