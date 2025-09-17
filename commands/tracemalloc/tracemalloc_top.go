package commands_tracemalloc

import "github.com/spf13/cobra"

func CreateCommandTraceMAllocTop() *cobra.Command {
	tracemalloc_top := &cobra.Command{
		Use:   "top",
		Short: "Show most common objects, the place that created them and their size.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return tracemalloc_top
}
