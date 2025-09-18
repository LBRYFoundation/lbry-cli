package rpc

import "context"
import "encoding/json"
import "fmt"
import "math/rand/v2"
import "os"
import "github.com/ybbus/jsonrpc/v3"

func ExecuteRPCCommand(method string, params ...interface{}) {
	rpcClient := jsonrpc.NewClient("http://localhost:5279/")

	request := jsonrpc.NewRequestWithID(rand.Int(), method, params...)
	response, err := rpcClient.CallRaw(context.Background(), request)

	if err != nil {
		fmt.Println("Could not connect to daemon. Are you sure it's running?")
		return
	}

	if response.Error != nil {
		PrintRawJSON(response.Error)
		os.Exit(1)
		return
	}
	PrintRawJSON(response.Result)
}

func PrintRawJSON(v any) {
	result, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(result))
}
