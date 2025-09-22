package commands_file

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFileList() *cobra.Command {
	file_list := &cobra.Command{
		Use:   "list",
		Short: "List files limited by optional filters",
		Run:   HandleCommandFileList,
	}

	file_list.Flags().String("sd_hash", "", "(str) get file with matching sd hash")
	file_list.Flags().String("file_name", "", "(str) get file with matching file name in the downloads folder")
	file_list.Flags().String("stream_hash", "", "(str) get file with matching stream hash")
	file_list.Flags().Int("rowid", -1, "(int) get file with matching row id")
	file_list.Flags().Int("added_on", -1, "(int) get file with matching time of insertion")
	file_list.Flags().String("claim_id", "", "(str) get file with matching claim id(s)")
	file_list.Flags().String("outpoint", "", "(str) get file with matching claim outpoint(s)")
	file_list.Flags().String("txid", "", "(str) get file with matching claim txid")
	file_list.Flags().Int("nout", -1, "(int) get file with matching claim nout")
	file_list.Flags().String("channel_claim_id", "", "(str) get file with matching channel claim id(s)")
	file_list.Flags().String("channel_name", "", "(str) get file with matching channel name")
	file_list.Flags().String("claim_name", "", "(str) get file with matching claim name")
	file_list.Flags().Int("blobs_in_stream", -1, "(int) get file with matching blobs in stream")
	file_list.Flags().String("download_path", "", "(str) get file with matching download path")
	file_list.Flags().Bool("uploading_to_reflector", false, "(bool) get files currently uploading to reflector")
	file_list.Flags().Bool("is_fully_reflected", false, "(bool) get files that have been uploaded to reflector")
	file_list.Flags().String("status", "", "(str) match by status, ( running | finished | stopped )")
	file_list.Flags().Bool("completed", false, "(bool) match only completed")
	file_list.Flags().Int("blobs_remaining", -1, "(int) amount of remaining blobs to download")
	file_list.Flags().String("sort", "", "(str) field to sort by (one of the above filter fields)")
	file_list.Flags().String("comparison", "", "(str) logical comparison, (eq | ne | g | ge | l | le | in)")
	file_list.Flags().Int("page", -1, "(int) page to return during paginating")
	file_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	file_list.Flags().String("wallet_id", "", "(str) add purchase receipts from this wallet")

	return file_list
}

func HandleCommandFileList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "rowid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "added_on")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "outpoint")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "blobs_in_stream")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "download_path")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "uploading_to_reflector")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_fully_reflected")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "status")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "completed")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "blobs_remaining")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sort")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "comparison")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("file_list", params)

}
