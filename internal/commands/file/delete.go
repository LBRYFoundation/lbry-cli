package commands_file

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFileDelete() *cobra.Command {
	file_delete := &cobra.Command{
		Use:   "file",
		Short: "Delete a LBRY file",
		Run:   HandleCommandFileDelete,
	}

	file_delete.Flags().Bool("delete_from_download_dir", false, "(bool) delete file from download directory, instead of just deleting blobs")
	file_delete.Flags().Bool("delete_all", false, "(bool) if there are multiple matching files, allow the deletion of multiple files. Otherwise do not delete anything.")
	file_delete.Flags().String("sd_hash", "", "(str) delete by file sd hash")
	file_delete.Flags().String("file_name", "", "(str) delete by file name in downloads folder")
	file_delete.Flags().String("stream_hash", "", "(str) delete by file stream hash")
	file_delete.Flags().Int("rowid", -1, "(int) delete by file row id")
	file_delete.Flags().String("claim_id", "", "(str) delete by file claim id")
	file_delete.Flags().String("txid", "", "(str) delete by file claim txid")
	file_delete.Flags().Int("nout", -1, "(int) delete by file claim nout")
	file_delete.Flags().String("claim_name", "", "(str) delete by file claim name")
	file_delete.Flags().String("channel_claim_id", "", "(str) delete by file channel claim id")
	file_delete.Flags().String("channel_name", "", "(str) delete by file channel claim name")

	return file_delete
}

func HandleCommandFileDelete(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "delete_from_download_dir")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "delete_all")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "rowid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "file_delete", params)
}
