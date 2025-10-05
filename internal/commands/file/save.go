package commands_file

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFileSave() *cobra.Command {
	file_save := &cobra.Command{
		Use:   "save",
		Short: "Start saving a file to disk.",
		Run:   HandleCommandFileSave,
	}

	file_save.Flags().String("file_name", "", "(str) file name to save to")
	file_save.Flags().String("download_directory", "", "(str) directory to save into")
	file_save.Flags().String("sd_hash", "", "(str) save file with matching sd hash")
	file_save.Flags().String("stream_hash", "", "(str) save file with matching stream hash")
	file_save.Flags().Int("rowid", -1, "(int) save file with matching row id")
	file_save.Flags().String("claim_id", "", "(str) save file with matching claim id")
	file_save.Flags().String("txid", "", "(str) save file with matching claim txid")
	file_save.Flags().Int("nout", -1, "(int) save file with matching claim nout")
	file_save.Flags().String("claim_name", "", "(str) save file with matching claim name")
	file_save.Flags().String("channel_claim_id", "", "(str) save file with matching channel claim id")
	file_save.Flags().String("channel_name", "", "(str) save file with matching channel claim name")

	return file_save
}

func HandleCommandFileSave(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "download_directory")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "file_save", params)
}
