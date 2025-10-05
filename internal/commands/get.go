package commands

import (
	"lbry/cli/internal/rpc"
	"strconv"

	"github.com/spf13/cobra"
)

func CreateCommandGet() *cobra.Command {
	get := &cobra.Command{
		Use:   "get",
		Short: "Download stream from a LBRY name.",
		Run:   HandleCommandGet,
	}

	//get.Flags().String("uri", "", "(str) uri of the content to download")
	get.Flags().String("file_name", "", "(str) specified name for the downloaded file, overrides the stream file name")
	get.Flags().String("download_directory", "", "(str) full path to the directory to download into")
	get.Flags().Int("timeout", -1, "(int) download timeout in number of seconds")
	get.Flags().Bool("save_file", false, "(bool) save the file to the downloads directory")
	get.Flags().String("wallet_id", "", "(str) wallet to check for claim purchase receipts")

	return get
}

func HandleCommandGet(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "uri")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "download_directory")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "timeout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "save_file")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["uri"]
		if exists {
			cmd.Help()
			return
		}
		params["uri"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["file_name"]
		if exists {
			cmd.Help()
			return
		}
		params["file_name"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["download_directory"]
		if exists {
			cmd.Help()
			return
		}
		params["download_directory"] = args[2]
	}
	if len(args) >= 4 {
		_, exists := params["timeout"]
		if exists {
			cmd.Help()
			return
		}
		val, _ := strconv.Atoi(args[3])
		params["timeout"] = val
	}
	if len(args) > 4 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, exists := params["uri"]
	if !exists {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "get", params)
}
