package commands_file

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFileReflect() *cobra.Command {
	file_reflect := &cobra.Command{
		Use:   "reflect",
		Short: "Reflect all the blobs in a file matching the filter criteria",
		Run:   HandleCommandFileReflect,
	}

	file_reflect.Flags().String("sd_hash", "", "(str) get file with matching sd hash")
	file_reflect.Flags().String("file_name", "", "(str) get file with matching file name in the downloads folder")
	file_reflect.Flags().String("stream_hash", "", "(str) get file with matching stream hash")
	file_reflect.Flags().Int("rowid", -1, "(int) get file with matching row id")
	file_reflect.Flags().String("reflector", "", "(str) reflector server, ip address or url by default choose a server from the config")

	return file_reflect
}

func HandleCommandFileReflect(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "stream_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "rowid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "reflector")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "file_reflect", params)
}
