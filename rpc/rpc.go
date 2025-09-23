package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ybbus/jsonrpc/v3"
)

func AddParameter[T any](parameters map[string]any, f *pflag.FlagSet, get func(name string) (T, error), name string) {
	if f.Changed(name) {
		value, _ := get(name)
		parameters[name] = value
	}
}

func GetURL(cmd *cobra.Command) string {
	if cmd.Flags().Changed("api") {
		api, _ := cmd.Flags().GetString("api")
		if !strings.HasPrefix(api, "//") && !regexp.MustCompile(`^[A-Za-z][A-Za-z0-9+.-]*://`).MatchString(api) {
			api = "//" + api
		}
		url, err := url.Parse(api)
		if err != nil {
			cmd.Help()
		}
		if url.Scheme == "" {
			// Use HTTP if scheme is absent
			url.Scheme = "http"
		}
		if url.Path == "" {
			// Use "/lbryapi" if path is absent
			url.Path = "/lbryapi"
		}
		return url.String()
	}
	return "http://localhost:5279/"
}

func ExecuteRPCCommand(url string, method string, params ...interface{}) {
	rpcClient := jsonrpc.NewClient(url)

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
