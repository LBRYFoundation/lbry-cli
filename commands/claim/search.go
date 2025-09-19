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
	name, _ := cmd.Flags().GetString("name")
	text, _ := cmd.Flags().GetString("text")
	claim_id, _ := cmd.Flags().GetString("claim_id")
	claim_ids, _ := cmd.Flags().GetStringArray("claim_ids")
	txid, _ := cmd.Flags().GetString("txid")
	nout, _ := cmd.Flags().GetString("nout")
	channel, _ := cmd.Flags().GetString("channel")
	channel_ids, _ := cmd.Flags().GetStringArray("channel_ids")
	not_channel_ids, _ := cmd.Flags().GetStringArray("not_channel_ids")
	has_channel_signature, _ := cmd.Flags().GetBool("has_channel_signature")
	valid_channel_signature, _ := cmd.Flags().GetBool("valid_channel_signature")
	invalid_channel_signature, _ := cmd.Flags().GetBool("invalid_channel_signature")
	limit_claims_per_channel, _ := cmd.Flags().GetInt("limit_claims_per_channel")
	is_controlling, _ := cmd.Flags().GetBool("is_controlling")
	public_key_id, _ := cmd.Flags().GetString("public_key_id")
	height, _ := cmd.Flags().GetInt("height")
	timestamp, _ := cmd.Flags().GetInt("timestamp")
	creation_height, _ := cmd.Flags().GetInt("creation_height")
	creation_timestamp, _ := cmd.Flags().GetInt("creation_timestamp")
	activation_height, _ := cmd.Flags().GetInt("activation_height")
	expiration_height, _ := cmd.Flags().GetInt("expiration_height")
	release_time, _ := cmd.Flags().GetInt("release_time")
	amount, _ := cmd.Flags().GetInt("amount")
	support_amount, _ := cmd.Flags().GetInt("support_amount")
	effective_amount, _ := cmd.Flags().GetInt("effective_amount")
	trending_score, _ := cmd.Flags().GetInt("trending_score")
	trending_group, _ := cmd.Flags().GetInt("trending_group")
	trending_mixed, _ := cmd.Flags().GetInt("trending_mixed")
	trending_local, _ := cmd.Flags().GetInt("trending_local")
	trending_global, _ := cmd.Flags().GetInt("trending_global")
	reposted_claim_id, _ := cmd.Flags().GetString("reposted_claim_id")
	reposted, _ := cmd.Flags().GetInt("reposted")
	claim_type, _ := cmd.Flags().GetString("claim_type")
	stream_types, _ := cmd.Flags().GetStringArray("stream_types")
	media_types, _ := cmd.Flags().GetStringArray("media_types")
	fee_currency, _ := cmd.Flags().GetString("fee_currency")
	fee_amount, _ := cmd.Flags().GetFloat64("fee_amount")
	duration, _ := cmd.Flags().GetInt("duration")
	any_tags, _ := cmd.Flags().GetStringArray("any_tags")
	all_tags, _ := cmd.Flags().GetStringArray("all_tags")
	not_tags, _ := cmd.Flags().GetStringArray("not_tags")
	any_languages, _ := cmd.Flags().GetStringArray("any_languages")
	all_languages, _ := cmd.Flags().GetStringArray("all_languages")
	not_languages, _ := cmd.Flags().GetStringArray("not_languages")
	any_locations, _ := cmd.Flags().GetStringArray("any_locations")
	all_locations, _ := cmd.Flags().GetStringArray("all_locations")
	not_locations, _ := cmd.Flags().GetStringArray("not_locations")
	page, _ := cmd.Flags().GetInt("page")
	page_size, _ := cmd.Flags().GetInt("page_size")
	order_by, _ := cmd.Flags().GetStringArray("order_by")
	no_totals, _ := cmd.Flags().GetBool("no_totals")
	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	include_purchase_receipt, _ := cmd.Flags().GetBool("include_purchase_receipt")
	include_is_my_output, _ := cmd.Flags().GetBool("include_is_my_output")
	remove_duplicates, _ := cmd.Flags().GetBool("remove_duplicates")
	has_source, _ := cmd.Flags().GetBool("has_source")
	sd_hash, _ := cmd.Flags().GetString("sd_hash")
	has_no_source, _ := cmd.Flags().GetBool("has_no_source")
	new_sdk_server, _ := cmd.Flags().GetString("new_sdk_server")

	// Check for arguments
	if len(args) >= 1 {
		name = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	params := map[string]any{}
	if name != "" {
		params["name"] = name
	}
	if text != "" {
		params["text"] = text
	}
	if claim_id != "" {
		params["claim_id"] = claim_id
	}
	if len(claim_ids) > 0 {
		params["claim_ids"] = claim_ids
	}
	if txid != "" {
		params["txid"] = txid
	}
	if nout != "" {
		params["nout"] = nout
	}
	if channel != "" {
		params["channel"] = channel
	}
	if len(channel_ids) > 0 {
		params["channel_ids"] = channel_ids
	}
	if len(not_channel_ids) > 0 {
		params["not_channel_ids"] = not_channel_ids
	}
	if has_channel_signature {
		params["has_channel_signature"] = has_channel_signature
	}
	if valid_channel_signature {
		params["valid_channel_signature"] = valid_channel_signature
	}
	if invalid_channel_signature {
		params["invalid_channel_signature"] = invalid_channel_signature
	}
	if limit_claims_per_channel != -1 {
		params["limit_claims_per_channel"] = limit_claims_per_channel
	}
	if is_controlling {
		params["is_controlling"] = is_controlling
	}
	if public_key_id != "" {
		params["public_key_id"] = public_key_id
	}
	if height != -1 {
		params["height"] = height
	}
	if timestamp != -1 {
		params["timestamp"] = timestamp
	}
	if creation_height != -1 {
		params["creation_height"] = creation_height
	}
	if creation_timestamp != -1 {
		params["creation_timestamp"] = creation_timestamp
	}
	if activation_height != -1 {
		params["activation_height"] = activation_height
	}
	if expiration_height != -1 {
		params["expiration_height"] = expiration_height
	}
	if release_time != -1 {
		params["release_time"] = release_time
	}
	if amount != -1 {
		params["amount"] = amount
	}
	if support_amount != -1 {
		params["support_amount"] = support_amount
	}
	if effective_amount != -1 {
		params["effective_amount"] = effective_amount
	}
	if trending_score != -1 {
		params["trending_score"] = trending_score
	}
	if trending_group != -1 {
		params["trending_group"] = trending_group
	}
	if trending_mixed != -1 {
		params["trending_mixed"] = trending_mixed
	}
	if trending_local != -1 {
		params["trending_local"] = trending_local
	}
	if trending_global != -1 {
		params["trending_global"] = trending_global
	}
	if reposted_claim_id != "" {
		params["reposted_claim_id"] = reposted_claim_id
	}
	if reposted != -1 {
		params["reposted"] = reposted
	}
	if claim_type != "" {
		params["claim_type"] = claim_type
	}
	if len(stream_types) > 0 {
		params["stream_types"] = stream_types
	}
	if len(media_types) > 0 {
		params["media_types"] = media_types
	}
	if fee_currency != "" {
		params["fee_currency"] = fee_currency
	}
	if fee_amount != -1 {
		params["fee_amount"] = fee_amount
	}
	if duration != -1 {
		params["duration"] = duration
	}
	if len(any_tags) > 0 {
		params["any_tags"] = any_tags
	}
	if len(all_tags) > 0 {
		params["all_tags"] = all_tags
	}
	if len(not_tags) > 0 {
		params["not_tags"] = not_tags
	}
	if len(any_languages) > 0 {
		params["any_languages"] = any_languages
	}
	if len(all_languages) > 0 {
		params["all_languages"] = all_languages
	}
	if len(not_languages) > 0 {
		params["not_languages"] = not_languages
	}
	if len(any_locations) > 0 {
		params["any_locations"] = any_locations
	}
	if len(all_locations) > 0 {
		params["all_locations"] = all_locations
	}
	if len(not_locations) > 0 {
		params["not_locations"] = not_locations
	}
	if page != -1 {
		params["page"] = page
	}
	if page_size != -1 {
		params["page_size"] = page_size
	}
	if len(order_by) > 0 {
		params["order_by"] = order_by
	}
	if no_totals {
		params["no_totals"] = no_totals
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if include_purchase_receipt {
		params["include_purchase_receipt"] = include_purchase_receipt
	}
	if include_is_my_output {
		params["include_is_my_output"] = include_is_my_output
	}
	if remove_duplicates {
		params["remove_duplicates"] = remove_duplicates
	}
	if has_source {
		params["has_source"] = has_source
	}
	if sd_hash != "" {
		params["sd_hash"] = sd_hash
	}
	if has_no_source {
		params["has_no_source"] = has_no_source
	}
	if new_sdk_server != "" {
		params["new_sdk_server"] = new_sdk_server
	}

	rpc.ExecuteRPCCommand("claim_search", params)
}
