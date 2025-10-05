package commands_blob

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandBlobReflect() *cobra.Command {
	blob_reflect := &cobra.Command{
		Use:   "reflect",
		Short: "Reflects specified blobs",
		Run:   HandleCommandBlobReflect,
	}

	//blob_reflect.Flags().StringArray("blob_hashes", nil, "")
	blob_reflect.Flags().String("reflector_server", "", "(str) reflector address")

	return blob_reflect
}

func HandleCommandBlobReflect(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "blob_hashes")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "reflector_server")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["blob_hashes"]
		if exists {
			cmd.Help()
			return
		}
		// Variable arguments
		params["blob_hashes"] = args
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "blob_reflect", params)
}
