package commands

import (
	"lbry/cli/rpc"

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
	publish.Flags().StringArray("tags", []string{}, "")
	publish.Flags().StringArray("languages", []string{}, "")
	publish.Flags().StringArray("locations", []string{}, "")
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
	publish.Flags().StringArray("funding_account_ids", []string{}, "")
	publish.Flags().String("claim_address", "", "")
	publish.Flags().Bool("preview", false, "")
	publish.Flags().Bool("blocking", false, "")

	return publish
}

func HandleCommandPublish(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	bid, _ := cmd.Flags().GetString("bid")
	cmd.Flags().GetString("file_path")
	cmd.Flags().GetString("file_name")
	cmd.Flags().GetString("file_hash")
	cmd.Flags().GetBool("validate_file")
	cmd.Flags().GetBool("optimize_file")
	cmd.Flags().GetString("fee_currency")
	cmd.Flags().GetFloat64("fee_amount")
	cmd.Flags().GetString("fee_address")
	cmd.Flags().GetString("title")
	cmd.Flags().GetString("description")
	cmd.Flags().GetString("author")
	cmd.Flags().GetStringArray("tags")
	cmd.Flags().GetStringArray("languages")
	cmd.Flags().GetStringArray("locations")
	cmd.Flags().GetString("license")
	cmd.Flags().GetString("license_url")
	cmd.Flags().GetString("thumbnail_url")
	cmd.Flags().GetInt("release_time")
	cmd.Flags().GetInt("width")
	cmd.Flags().GetInt("height")
	cmd.Flags().GetInt("duration")
	cmd.Flags().GetString("sd_hash")
	cmd.Flags().GetString("channel_id")
	cmd.Flags().GetString("channel_name")
	cmd.Flags().GetString("channel_account_id")
	cmd.Flags().GetString("account_id")
	cmd.Flags().GetString("wallet_id")
	cmd.Flags().GetStringArray("funding_account_ids")
	cmd.Flags().GetString("claim_address")
	cmd.Flags().GetBool("preview")
	cmd.Flags().GetBool("blocking")

	// Check for arguments
	if len(args) >= 1 {
		name = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	if name == "" {
		cmd.Help()
		return
	}

	// Create parameter map
	params := map[string]any{
		"name": name,
	}
	if bid != "" {
		params["bid"] = bid
	}

	rpc.ExecuteRPCCommand("publish", params)
}
