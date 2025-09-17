package commands

import "context"
import "encoding/json"
import "fmt"
import "github.com/spf13/cobra"
import "github.com/ybbus/jsonrpc/v3"

func CreateCommandVersion() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Get lbrynet API server version information",
		Run:   HandleCommandVersion,
	}

	return version
}

func HandleCommandVersion(cmd *cobra.Command, args []string) {
	rpcClient := jsonrpc.NewClient("http://localhost:5279/")
	resp, err := rpcClient.Call(context.Background(), "version")
	if err != nil {
		fmt.Println("Could not connect to daemon. Are you sure it's running?")
		return
	}
	result, _ := json.MarshalIndent(resp.Result, "", "\t")
	fmt.Println(string(result))

}
