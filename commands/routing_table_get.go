package commands

import "context"
import "encoding/json"
import "fmt"
import "github.com/spf13/cobra"
import "github.com/ybbus/jsonrpc/v3"

func CreateCommandRoutingTableGet() *cobra.Command {
	routing_table_get := &cobra.Command{
		Use:   "routing_table_get",
		Short: "Get DHT routing information",
		Run:   HandleCommandRoutingTableGet,
	}

	return routing_table_get
}

func HandleCommandRoutingTableGet(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}
	rpcClient := jsonrpc.NewClient("http://localhost:5279/")
	resp, err := rpcClient.Call(context.Background(), "routing_table_get")
	if err != nil {
		fmt.Println("Could not connect to daemon. Are you sure it's running?")
		return
	}
	result, _ := json.MarshalIndent(resp.Result, "", "\t")
	fmt.Println(string(result))
}
