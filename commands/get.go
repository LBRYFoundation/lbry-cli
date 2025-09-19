package commands

import (
	"lbry/cli/rpc"
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
	uri := ""
	//uri, _ := cmd.Flags().GetString("uri")

	file_name, _ := cmd.Flags().GetString("file_name")
	download_directory, _ := cmd.Flags().GetString("download_directory")
	timeout, _ := cmd.Flags().GetInt("timeout")
	save_file, _ := cmd.Flags().GetBool("save_file")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		uri = args[0]
	}
	if len(args) >= 2 {
		file_name = args[1]
	}
	if len(args) >= 3 {
		download_directory = args[2]
	}
	if len(args) >= 4 {
		val, _ := strconv.Atoi(args[3])
		timeout = val
	}

	if uri == "" {
		cmd.Help()
		return
	}

	params := map[string]any{
		"uri": uri,
	}
	if file_name != "" {
		params["file_name"] = file_name
	}
	if download_directory != "" {
		params["download_directory"] = download_directory
	}
	if timeout != -1 {
		params["timeout"] = timeout
	}
	if save_file {
		params["save_file"] = save_file
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}

	rpc.ExecuteRPCCommand("get", params)
}
