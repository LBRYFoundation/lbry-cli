package commands_claim

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandClaimSearch() *cobra.Command {
	claim_search := &cobra.Command{
		Use:   "search",
		Short: "Search for stream and channel claims on the blockchain.",
		Run:   HandleCommandClaimSearch,
	}

	claim_search.Flags().String("name", "", "(str) claim name (normalized)")
	claim_search.Flags().String("text", "", "(str) full text search")
	claim_search.Flags().String("claim_id", "", "(str) full or partial claim id")
	claim_search.Flags().StringArray("claim_ids", nil, "(list) list of full claim ids")
	claim_search.Flags().String("txid", "", "(str) transaction id")
	claim_search.Flags().String("nout", "", "(str) position in the transaction")
	claim_search.Flags().String("channel", "", "(str) claims signed by this channel (argument is a URL which automatically gets resolved), see --channel_ids if you need to filter by multiple channels at the same time, includes claims with invalid signatures, use in conjunction with --valid_channel_signature")
	claim_search.Flags().StringArray("channel_ids", nil, "(list) claims signed by any of these channels (arguments must be claim ids of the channels), includes claims with invalid signatures, implies --has_channel_signature, use in conjunction with --valid_channel_signature")
	claim_search.Flags().StringArray("not_channel_ids", nil, "(list) exclude claims signed by any of these channels (arguments must be claim ids of the channels)")
	claim_search.Flags().Bool("has_channel_signature", false, "(bool) claims with a channel signature (valid or invalid)")
	claim_search.Flags().Bool("valid_channel_signature", false, "(bool) claims with a valid channel signature or no signature, use in conjunction with --has_channel_signature to only get claims with valid signatures")
	claim_search.Flags().Bool("invalid_channel_signature", false, "(bool) claims with invalid channel signature or no signature, use in conjunction with --has_channel_signature to only get claims with invalid signatures")
	claim_search.Flags().Int("limit_claims_per_channel", -1, "(int) only return up to the specified number of claims per channel")
	claim_search.Flags().Bool("is_controlling", false, "(bool) winning claims of their respective name")
	claim_search.Flags().String("public_key_id", "", "(str) only return channels having this public key id, this is the same key as used in the wallet file to map channel certificate private keys: {'public_key_id': 'private key'}")
	claim_search.Flags().Int("height", -1, "(int) last updated block height (supports equality constraints)")
	claim_search.Flags().Int("timestamp", -1, "(int) last updated timestamp (supports equality constraints)")
	claim_search.Flags().Int("creation_height", -1, "(int) created at block height (supports equality constraints)")
	claim_search.Flags().Int("creation_timestamp", -1, "(int) created at timestamp (supports equality constraints)")
	claim_search.Flags().Int("activation_height", -1, "(int) height at which claim starts competing for name (supports equality constraints)")
	claim_search.Flags().Int("expiration_height", -1, "(int) height at which claim will expire (supports equality constraints)")
	claim_search.Flags().Int("release_time", -1, "(int) limit to claims self-described as having been released to the public on or after this UTC timestamp, when claim does not provide a release time the publish time is used instead (supports equality constraints)")
	claim_search.Flags().Int("amount", -1, "(int) limit by claim value (supports equality constraints)")
	claim_search.Flags().Int("support_amount", -1, "(int) limit by supports and tips received (supports equality constraints)")
	claim_search.Flags().Int("effective_amount", -1, "(int) limit by total value (initial claim value plus all tips and supports received), this amount is blank until claim has reached activation height (supports equality constraints)")
	claim_search.Flags().Int("trending_score", -1, "(int) limit by trending score (supports equality constraints)")
	claim_search.Flags().Int("trending_group", -1, "(int) DEPRECATED - instead please use trending_score")
	claim_search.Flags().Int("trending_mixed", -1, "(int) DEPRECATED - instead please use trending_score")
	claim_search.Flags().Int("trending_local", -1, "(int) DEPRECATED - instead please use trending_score")
	claim_search.Flags().Int("trending_global", -1, "(int) DEPRECATED - instead please use trending_score")
	claim_search.Flags().String("reposted_claim_id", "", "(str) all reposts of the specified original claim id")
	claim_search.Flags().Int("reposted", -1, "(int) claims reposted this many times (supports equality constraints)")
	claim_search.Flags().String("claim_type", "", "(str) filter by 'channel', 'stream', 'repost' or 'collection'")
	claim_search.Flags().StringArray("stream_types", nil, "(list) filter by 'video', 'image', 'document', etc")
	claim_search.Flags().StringArray("media_types", nil, "(list) filter by 'video/mp4', 'image/png', etc")
	claim_search.Flags().String("fee_currency", "", "(string) specify fee currency: LBC, BTC, USD")
	claim_search.Flags().Float64("fee_amount", -1, "(decimal) content download fee (supports equality constraints)")
	claim_search.Flags().Int("duration", -1, "(int) duration of video or audio in seconds (supports equality constraints)")
	claim_search.Flags().StringArray("any_tags", nil, "(list) find claims containing any of the tags")
	claim_search.Flags().StringArray("all_tags", nil, "(list) find claims containing every tag")
	claim_search.Flags().StringArray("not_tags", nil, "(list) find claims not containing any of these tags")
	claim_search.Flags().StringArray("any_languages", nil, "(list) find claims containing any of the languages")
	claim_search.Flags().StringArray("all_languages", nil, "(list) find claims containing every language")
	claim_search.Flags().StringArray("not_languages", nil, "(list) find claims not containing any of these languages")
	claim_search.Flags().StringArray("any_locations", nil, "(list) find claims containing any of the locations")
	claim_search.Flags().StringArray("all_locations", nil, "(list) find claims containing every location")
	claim_search.Flags().StringArray("not_locations", nil, "(list) find claims not containing any of these locations")
	claim_search.Flags().Int("page", -1, "(int) page to return during paginating")
	claim_search.Flags().Int("page_size", -1, "(int) number of items on page during pagination")
	claim_search.Flags().StringArray("order_by", nil, "(list) field to order by, default is descending order, to do an ascending order prepend ^ to the field name, eg. '^amount' available fields: 'name', 'height', 'release_time', 'publish_time', 'amount', 'effective_amount', 'support_amount', 'trending_group', 'trending_mixed', 'trending_local', 'trending_global', 'activation_height'")
	claim_search.Flags().Bool("no_totals", false, "(bool) do not calculate the total number of pages and items in result set (significant performance boost)")
	claim_search.Flags().String("wallet_id", "", "(str) wallet to check for claim purchase receipts")
	claim_search.Flags().Bool("include_purchase_receipt", false, "(bool) lookup and include a receipt if this wallet has purchased the claim")
	claim_search.Flags().Bool("include_is_my_output", false, "(bool) lookup and include a boolean indicating if claim being resolved is yours")
	claim_search.Flags().Bool("remove_duplicates", false, "(bool) removes duplicated content from search by picking either the original claim or the oldest matching repost")
	claim_search.Flags().Bool("has_source", false, "(bool) find claims containing a source field")
	claim_search.Flags().String("sd_hash", "", "(str)  find claims where the source stream descriptor hash matches (partially or completely) the given hexadecimal string")
	claim_search.Flags().Bool("has_no_source", false, "(bool) find claims not containing a source field")
	claim_search.Flags().String("new_sdk_server", "", "(str) URL of the new SDK server (EXPERIMENTAL)")

	return claim_search
}

func HandleCommandClaimSearch(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "text")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "txid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "nout")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "channel_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_channel_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "has_channel_signature")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "valid_channel_signature")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "invalid_channel_signature")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "limit_claims_per_channel")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "is_controlling")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "public_key_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "height")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "timestamp")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "creation_height")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "creation_timestamp")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "activation_height")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "expiration_height")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "release_time")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "support_amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "effective_amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "trending_score")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "trending_group")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "trending_mixed")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "trending_local")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "trending_global")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "reposted_claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "reposted")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_type")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "stream_types")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "media_types")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "fee_currency")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "fee_amount")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "duration")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "any_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "all_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "any_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "all_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "any_locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "all_locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "not_locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "order_by")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "no_totals")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_purchase_receipt")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_is_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "remove_duplicates")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "has_source")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "sd_hash")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "has_no_source")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "new_sdk_server")

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

	rpc.ExecuteRPCCommand("claim_search", params)
}
