package commands_stream

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStreamUpdate() *cobra.Command {
	stream_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing stream claim and if a new file is provided announce it to lbrynet.",
		Run:   HandleCommandStreamUpdate,
	}

	stream_update.Flags().String("claim_id", "", "(str) id of the stream claim to update")
	stream_update.Flags().Float64("bid", -1, "(decimal) amount to back the claim")
	stream_update.Flags().String("file_path", "", "(str) path to file to be associated with name")
	stream_update.Flags().Bool("validate_file", false, "(bool) validate that the video container and encodings match common web browser support or that optimization succeeds if specified. FFmpeg is required and file_path must be specified.")
	stream_update.Flags().Bool("optimize_file", false, "(bool) transcode the video & audio if necessary to ensure common web browser support. FFmpeg is required and file_path must be specified.")
	stream_update.Flags().String("file_name", "", "(str) override file name, defaults to name from file_path.")
	stream_update.Flags().String("file_size", "", "(str) override file size, otherwise automatically computed.")
	stream_update.Flags().String("file_hash", "", "(str) override file hash, otherwise automatically computed.")
	stream_update.Flags().String("fee_currency", "", "(string) specify fee currency")
	stream_update.Flags().Float64("fee_amount", -1, "(decimal) content download fee")
	stream_update.Flags().String("fee_address", "", "(str) address where to send fee payments, will use value from --claim_address if not provided")
	stream_update.Flags().Bool("clear_fee", false, "(bool) clear previously set fee")
	stream_update.Flags().String("title", "", "(str) title of the publication")
	stream_update.Flags().String("description", "", "(str) description of the publication")
	stream_update.Flags().String("author", "", "(str) author of the publication. The usage for this field is not the same as for channels. The author field is used to credit an author who is not the publisher and is not represented by the channel. For example, a pdf file of 'The Odyssey' has an author of 'Homer' but may by published to a channel such as '@classics', or to no channel at all")
	stream_update.Flags().StringArray("tags", nil, "(list) add content tags")
	stream_update.Flags().Bool("clear_tags", false, "(bool) clear existing tags (prior to adding new ones)")
	stream_update.Flags().StringArray("languages", nil, "(list) languages used by the channel, using RFC 5646 format, eg:\n\tfor English `--languages=en`\n\tfor Spanish (Spain) `--languages=es-ES`\n\tfor Spanish (Mexican) `--languages=es-MX`\n\tfor Chinese (Simplified) `--languages=zh-Hans`\n\tfor Chinese (Traditional) `--languages=zh-Hant`")
	stream_update.Flags().Bool("clear_languages", false, "(bool) clear existing languages (prior to adding new ones)")
	stream_update.Flags().StringArray("locations", nil, "(list) locations of the stream, consisting of 2 letter `country` code and a `state`, `city` and a postal `code` along with a `latitude` and `longitude`.\nfor JSON RPC: pass a dictionary with aforementioned attributes as keys, eg:\n\t...\n\t\"locations\": [{'country': 'US', 'state': 'NH'}]\n\t...\nfor command line: pass a colon delimited list with values in the following order:\n\n\t\"COUNTRY:STATE:CITY:CODE:LATITUDE:LONGITUDE\"\n\nmaking sure to include colon for blank values, for example to provide only the city:\n\n\t... --locations=\"::Manchester\"\n\nwith all values set:\n\n\t... --locations=\"US:NH:Manchester:03101:42.990605:-71.460989\"\n\noptionally, you can just pass the \"LATITUDE:LONGITUDE\":\n\n\t... --locations=\"42.990605:-71.460989\"\n\nfinally, you can also pass JSON string of dictionary on the command line as you would via JSON RPC\n\n\t... --locations=\"{'country': 'US', 'state': 'NH'}\"")
	stream_update.Flags().Bool("clear_locations", false, "(bool) clear existing locations (prior to adding new ones)")
	stream_update.Flags().String("license", "", "(str) publication license")
	stream_update.Flags().String("license_url", "", "(str) publication license url")
	stream_update.Flags().String("thumbnail_url", "", "(str) thumbnail url")
	stream_update.Flags().Int("release_time", -1, "(int) original public release of content, seconds since UNIX epoch")
	stream_update.Flags().Int("width", -1, "(int) image/video width, automatically calculated from media file")
	stream_update.Flags().Int("height", -1, "(int) image/video height, automatically calculated from media file")
	stream_update.Flags().Int("duration", -1, "(int) audio/video duration in seconds, automatically calculated")
	stream_update.Flags().String("sd_hash", "", "(str) sd_hash of stream")
	stream_update.Flags().String("channel_id", "", "(str) claim id of the publisher channel")
	stream_update.Flags().String("channel_name", "", "(str) name of the publisher channel")
	stream_update.Flags().Bool("clear_channel", false, "(bool) remove channel signature")
	stream_update.Flags().String("channel_account_id", "", "(str) one or more account ids for accounts to look in for channel certificates, defaults to all accounts.")
	stream_update.Flags().String("account_id", "", "(str) account in which to look for stream (default: all)")
	stream_update.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	stream_update.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	stream_update.Flags().String("claim_address", "", "(str) address where the claim is sent to, if not specified it will be determined automatically from the account")
	stream_update.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	stream_update.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")
	stream_update.Flags().Bool("replace", false, "(bool) instead of modifying specific values on the stream, this will clear all existing values and only save passed in values, useful for form submissions where all values are always set")

	return stream_update
}

func HandleCommandStreamUpdate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_path")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "validate_file")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "optimize_file")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "file_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "fee_currency")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "fee_amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "fee_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_fee")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "author")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_locations")
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
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_channel")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "blocking")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "replace")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["claim_id"]
		if exists {
			cmd.Help()
			return
		}
		params["claim_id"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "stream_update")
}
