package commands_tracemalloc

import (
	"lbry/cli/rpc"
	"strconv"

	"github.com/spf13/cobra"
)

func CreateCommandTraceMAllocTop() *cobra.Command {
	tracemalloc_top := &cobra.Command{
		Use:   "top",
		Short: "Show most common objects, the place that created them and their size.",
		Run:   HandleCommandTraceMAllocTop,
	}

	tracemalloc_top.Flags().Int("items", -1, "(int) maximum items to return, from the most common")

	return tracemalloc_top
}

func HandleCommandTraceMAllocTop(cmd *cobra.Command, args []string) {
	items, _ := cmd.Flags().GetInt("items")

	// Check for arguments
	if len(args) >= 1 {
		val, _ := strconv.Atoi(args[0])
		items = val
	}

	params := map[string]any{}
	if items != -1 {
		params["items"] = items
	}

	rpc.ExecuteRPCCommand("tracemalloc_top", params)
}
