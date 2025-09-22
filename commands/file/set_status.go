package commands_file

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFileSetStatus() *cobra.Command {
	file_set_status := &cobra.Command{
		Use:   "set_status",
		Short: "Start or stop downloading a file",
		Run:   HandleCommandFileSetStatus,
	}

	file_set_status.Flags().String("status", "", "(str) one of \"start\" or \"stop\"")
	file_set_status.Flags().String("sd_hash", "", "(str) set status of file with matching sd hash")
	file_set_status.Flags().String("file_name", "", "(str) set status of file with matching file name in the downloads folder")
	file_set_status.Flags().String("stream_hash", "", "(str) set status of file with matching stream hash")
	file_set_status.Flags().Int("rowid", -1, "(int) set status of file with matching row id")

	return file_set_status
}

func HandleCommandFileSetStatus(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "status")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "rowid")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["status"]
		if exists {
			cmd.Help()
			return
		}
		params["status"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("file_set_status", params)
}
