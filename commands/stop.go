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
	rpcClient := jsonrpc.NewClient("http://localhost:5279/")
	resp, err := rpcClient.Call(context.Background(), "stop")
	if err != nil {
		fmt.Println("Could not connect to daemon. Are you sure it's running?")
		return
	}
	result, _ := json.MarshalIndent(resp.Result, "", "\t")
	fmt.Println(string(result))
}
