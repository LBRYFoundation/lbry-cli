package commands

import "context"
import "encoding/json"
import "fmt"
import "github.com/spf13/cobra"
import "github.com/ybbus/jsonrpc/v3"

func CreateCommandFFmpegFind() *cobra.Command {
	ffmpeg_find := &cobra.Command{
		Use:   "ffmpeg_find",
		Short: "Get ffmpeg installation information",
		Run:   HandleCommandFFmpegFind,
	}

	return ffmpeg_find
}

func HandleCommandFFmpegFind(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}
	rpcClient := jsonrpc.NewClient("http://localhost:5279/")
	resp, err := rpcClient.Call(context.Background(), "ffmpeg_find")
	if err != nil {
		fmt.Println("Could not connect to daemon. Are you sure it's running?")
		return
	}
	result, _ := json.MarshalIndent(resp.Result, "", "\t")
	fmt.Println(string(result))
}
