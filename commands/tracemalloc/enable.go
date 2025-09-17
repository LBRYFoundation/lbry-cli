package commands_tracemalloc

import "github.com/spf13/cobra"

func CreateCommandTraceMAllocEnable() *cobra.Command {
	tracemalloc_enable := &cobra.Command{
		Use:   "enable",
		Short: "Enable tracemalloc memory tracing",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return tracemalloc_enable
}
