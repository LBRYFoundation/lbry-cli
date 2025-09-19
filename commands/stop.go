package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ybbus/jsonrpc/v3"
)

func CreateCommandStop() *cobra.Command {
	stop := &cobra.Command{
		Use:   "stop",
		Short: "Stop lbrynet API server.",
		Run:   HandleCommandStop,
	}

	return stop
}

func HandleCommandStop(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("stop")
}
