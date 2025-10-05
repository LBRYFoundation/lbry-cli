package commands_tracemalloc

import (
	"lbry/cli/internal/rpc"
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
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "items")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["password"]
		if exists {
			cmd.Help()
			return
		}
		val, _ := strconv.Atoi(args[0])
		params["password"] = val
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "tracemalloc_top", params)
}
