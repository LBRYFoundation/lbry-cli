package commands_collection

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandCollectionUpdate() *cobra.Command {
	collection_update := &cobra.Command{
		Use:   "update",
		Short: "Update an existing collection claim.",
		Run:   HandleCommandCollectionUpdate,
	}

	collection_update.Flags().String("claim_id", "", "(str) claim_id of the collection to update")
	collection_update.Flags().Float64("bid", -1, "(decimal) amount to back the claim")
	collection_update.Flags().StringArray("claims", nil, "(list) claim ids")
	collection_update.Flags().StringArray("clear_claims", nil, "(bool) clear existing claim references (prior to adding new ones)")
	collection_update.Flags().String("title", "", "(str) title of the collection")
	collection_update.Flags().String("description", "", "(str) description of the collection")
	collection_update.Flags().StringArray("tags", nil, "(list) add content tags")
	collection_update.Flags().Bool("clear_tags", false, "(bool) clear existing tags (prior to adding new ones)")
	collection_update.Flags().StringArray("languages", nil, "(list) languages used by the collection, using RFC 5646 format, eg:\nfor English `--languages=en`\nfor Spanish (Spain) `--languages=es-ES`\nfor Spanish (Mexican) `--languages=es-MX`\nfor Chinese (Simplified) `--languages=zh-Hans`\nfor Chinese (Traditional) `--languages=zh-Hant`")
	collection_update.Flags().Bool("clear_languages", false, "(bool) clear existing languages (prior to adding new ones)")
	collection_update.Flags().StringArray("locations", nil, "(list) locations of the collection, consisting of 2 letter `country` code and a `state`, `city` and a postal `code` along with a `latitude` and `longitude`.\nfor JSON RPC: pass a dictionary with aforementioned attributes as keys, eg:\n\t...\n\t\"locations\": [{'country': 'US', 'state': 'NH'}]\n\t...\nfor command line: pass a colon delimited list with values in the following order:\n\n\t\"COUNTRY:STATE:CITY:CODE:LATITUDE:LONGITUDE\"\n\nmaking sure to include colon for blank values, for example to provide only the city:\n\n\t... --locations=\"::Manchester\"\n\nwith all values set:\n\n\t... --locations=\"US:NH:Manchester:03101:42.990605:-71.460989\"\n\noptionally, you can just pass the \"LATITUDE:LONGITUDE\":\n\n\t... --locations=\"42.990605:-71.460989\"\n\nfinally, you can also pass JSON string of dictionary on the command line as you would via JSON RPC\n\n\t... --locations=\"{'country': 'US', 'state': 'NH'}\"")
	collection_update.Flags().Bool("clear_locations", false, "(bool) clear existing locations (prior to adding new ones)")
	collection_update.Flags().String("thumbnail_url", "", "(str) thumbnail url")
	collection_update.Flags().String("channel_id", "", "(str) claim id of the publisher channel")
	collection_update.Flags().String("channel_name", "", "(str) name of the publisher channel")
	collection_update.Flags().String("channel_account_id", "", "(str) one or more account ids for accounts to look in for channel certificates, defaults to all accounts.")
	collection_update.Flags().String("account_id", "", "(str) account in which to look for collection (default: all)")
	collection_update.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	collection_update.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	collection_update.Flags().String("claim_address", "", "(str) address where the collection is sent")
	collection_update.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	collection_update.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")
	collection_update.Flags().Bool("replace", false, "(bool) instead of modifying specific values on the collection, this will clear all existing values and only save passed in values, useful for form submissions where all values are always set")

	return collection_update
}

func HandleCommandCollectionUpdate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claims")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "clear_claims")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "clear_locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "thumbnail_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "collection_update", params)
}
