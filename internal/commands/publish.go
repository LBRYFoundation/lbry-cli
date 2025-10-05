package commands

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandPublish() *cobra.Command {
	publish := &cobra.Command{
		Use:   "publish",
		Short: "Create or replace a stream claim at a given name (use 'stream create/update' for more control).",
		Run:   HandleCommandPublish,
	}

	publish.Flags().String("name", "", "")
	publish.Flags().String("bid", "", "")
	publish.Flags().String("file_path", "", "")
	publish.Flags().String("file_name", "", "")
	publish.Flags().String("file_hash", "", "")
	publish.Flags().Bool("validate_file", false, "")
	publish.Flags().Bool("optimize_file", false, "")
	publish.Flags().String("fee_currency", "", "")
	publish.Flags().Float64("fee_amount", -1, "")
	publish.Flags().String("fee_address", "", "")
	publish.Flags().String("title", "", "")
	publish.Flags().String("description", "", "")
	publish.Flags().String("author", "", "")
	publish.Flags().StringArray("tags", nil, "")
	publish.Flags().StringArray("languages", nil, "")
	publish.Flags().StringArray("locations", nil, "")
	publish.Flags().String("license", "", "")
	publish.Flags().String("license_url", "", "")
	publish.Flags().String("thumbnail_url", "", "")
	publish.Flags().Int("release_time", -1, "")
	publish.Flags().Int("width", -1, "")
	publish.Flags().Int("height", -1, "")
	publish.Flags().Int("duration", -1, "")
	publish.Flags().String("sd_hash", "", "")
	publish.Flags().String("channel_id", "", "")
	publish.Flags().String("channel_name", "", "")
	publish.Flags().String("channel_account_id", "", "")
	publish.Flags().String("account_id", "", "")
	publish.Flags().String("wallet_id", "", "")
	publish.Flags().StringArray("funding_account_ids", nil, "")
	publish.Flags().String("claim_address", "", "")
	publish.Flags().Bool("preview", false, "")
	publish.Flags().Bool("blocking", false, "")

	return publish
}

func HandleCommandPublish(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_path")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "validate_file")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "optimize_file")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "fee_currency")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "fee_amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "fee_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "author")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "license")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "license_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "thumbnail_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "release_time")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "width")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "height")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "duration")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["name"]
		if exists {
			cmd.Help()
			return
		}
		params["name"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, exists := params["name"]
	if !exists {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "publish", params)
}
