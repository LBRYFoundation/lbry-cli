package commands_channel

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandChannelUpdate() *cobra.Command {
	channel_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing channel claim.",
		Run:   HandleCommandChannelUpdate,
	}

	channel_update.Flags().String("claim_id", "", "(str) claim_id of the channel to update")
	channel_update.Flags().Float64("bid", -1, "(decimal) amount to back the claim")
	channel_update.Flags().String("title", "", "(str) title of the publication")
	channel_update.Flags().String("description", "", "(str) description of the publication")
	channel_update.Flags().String("email", "", "(str) email of channel owner")
	channel_update.Flags().String("website_url", "", "(str) website url")
	channel_update.Flags().StringArray("featured", nil, "(list) claim_ids of featured content in channel")
	channel_update.Flags().Bool("clear_featured", false, "(bool) clear existing featured content (prior to adding new ones)")
	channel_update.Flags().StringArray("tags", nil, "(list) content tags")
	channel_update.Flags().Bool("clear_tags", false, "(bool) clear existing tags (prior to adding new ones)")
	channel_update.Flags().StringArray("languages", nil, "(list) languages used by the channel, using RFC 5646 format, eg:\nfor English `--languages=en`\nfor Spanish (Spain) `--languages=es-ES`\nfor Spanish (Mexican) `--languages=es-MX`\nfor Chinese (Simplified) `--languages=zh-Hans`\nfor Chinese (Traditional) `--languages=zh-Hant`")
	channel_update.Flags().Bool("clear_languages", false, "(bool) clear existing languages (prior to adding new ones)")
	channel_update.Flags().StringArray("locations", nil, "(list) locations of the channel, consisting of 2 letter `country` code and a `state`, `city` and a postal `code` along with a `latitude` and `longitude`.\nfor JSON RPC: pass a dictionary with aforementioned attributes as keys, eg:\n\t...\n\t\"locations\": [{'country': 'US', 'state': 'NH'}]\n\t...\nfor command line: pass a colon delimited list with values in the following order:\n\n\t\"COUNTRY:STATE:CITY:CODE:LATITUDE:LONGITUDE\"\n\nmaking sure to include colon for blank values, for example to provide only the city:\n\n\t... --locations=\"::Manchester\"\n\nwith all values set:\n\n\t... --locations=\"US:NH:Manchester:03101:42.990605:-71.460989\"\n\noptionally, you can just pass the \"LATITUDE:LONGITUDE\":\n\n\t... --locations=\"42.990605:-71.460989\"\n\nfinally, you can also pass JSON string of dictionary on the command line as you would via JSON RPC\n\n\t... --locations=\"{'country': 'US', 'state': 'NH'}\"")
	channel_update.Flags().Bool("clear_locations", false, "(bool) clear existing locations (prior to adding new ones)")
	channel_update.Flags().String("thumbnail_url", "", "(str) thumbnail url")
	channel_update.Flags().String("cover_url", "", "(str) url of cover image")
	channel_update.Flags().String("account_id", "", "(str) account in which to look for channel (default: all)")
	channel_update.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	channel_update.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	channel_update.Flags().String("claim_address", "", "(str) address where the channel is sent")
	channel_update.Flags().Bool("new_signing_key", false, "(bool) generate a new signing key, will invalidate all previous publishes")
	channel_update.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	channel_update.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")
	channel_update.Flags().Bool("replace", false, "(bool) instead of modifying specific values on the channel, this will clear all existing values and only save passed in values, useful for form submissions where all values are always set")

	return channel_update
}

func HandleCommandChannelUpdate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "email")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "website_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "featured")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_featured")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "thumbnail_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "cover_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "new_signing_key")
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
	if len(args) >= 2 {
		_, exists := params["bid"]
		if exists {
			cmd.Help()
			return
		}
		params["bid"] = args[1]
	}
	if len(args) > 2 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "channel_update", params)
}
