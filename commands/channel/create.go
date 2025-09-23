package commands_channel

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandChannelCreate() *cobra.Command {
	channel_create := &cobra.Command{
		Use:   "create",
		Short: "Create a new channel by generating a channel private key and establishing an '@' prefixed claim.",
		Run:   HandleCommandChannelCreate,
	}

	channel_create.Flags().String("name", "", "(str) name of the channel prefixed with '@'")
	channel_create.Flags().Float64("bid", -1, "(decimal) amount to back the claim")
	channel_create.Flags().Bool("allow_duplicate_name", false, "(bool) create new channel even if one already exists with given name. default: false.")
	channel_create.Flags().String("title", "", "(str) title of the publication")
	channel_create.Flags().String("description", "", "(str) description of the publication")
	channel_create.Flags().String("email", "", "(str) email of channel owner")
	channel_create.Flags().String("website_url", "", "(str) website url")
	channel_create.Flags().StringArray("featured", nil, "(list) claim_ids of featured content in channel")
	channel_create.Flags().StringArray("tags", nil, "(list) content tags")
	channel_create.Flags().StringArray("languages", nil, "(list) languages used by the channel, using RFC 5646 format, eg:\n\tfor English `--languages=en`\n\tfor Spanish (Spain) `--languages=es-ES`\n\tfor Spanish (Mexican) `--languages=es-MX`\n\tfor Chinese (Simplified) `--languages=zh-Hans`\n\tfor Chinese (Traditional) `--languages=zh-Hant`")
	channel_create.Flags().StringArray("locations", nil, "(list) locations of the channel, consisting of 2 letter `country` code and a `state`, `city` and a postal `code` along with a `latitude` and `longitude`.\nfor JSON RPC: pass a dictionary with aforementioned attributes as keys, eg:\n\t...\n\t\"locations\": [{'country': 'US', 'state': 'NH'}]\n\t...\nfor command line: pass a colon delimited list with values in the following order:\n\n\t\"COUNTRY:STATE:CITY:CODE:LATITUDE:LONGITUDE\"\n\nmaking sure to include colon for blank values, for example to provide only the city:\n\n\t... --locations=\"::Manchester\"\n\nwith all values set:\n\n\t... --locations=\"US:NH:Manchester:03101:42.990605:-71.460989\"\n\noptionally, you can just pass the \"LATITUDE:LONGITUDE\":\n\n\t... --locations=\"42.990605:-71.460989\"\n\nfinally, you can also pass JSON string of dictionary on the command line as you would via JSON RPC\n\n\t... --locations=\"{'country': 'US', 'state': 'NH'}\"")
	channel_create.Flags().String("thumbnail_url", "", "(str) thumbnail url")
	channel_create.Flags().String("cover_url", "", "(str) url of cover image")
	channel_create.Flags().String("account_id", "", "(str) account to use for holding the transaction")
	channel_create.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	channel_create.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	channel_create.Flags().StringArray("claim_address", nil, "(str) address where the channel is sent to, if not specified it will be determined automatically from the account")
	channel_create.Flags().StringArray("preview", nil, "(bool) do not broadcast the transaction")
	channel_create.Flags().StringArray("blocking", nil, "(bool) wait until transaction is in mempool")

	return channel_create
}

func HandleCommandChannelCreate(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "allow_duplicate_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "email")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "website_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "featured")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "languages")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "locations")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "thumbnail_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "cover_url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "funding_account_ids")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "claim_address")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "preview")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "preview")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["name"]
		if exists {
			cmd.Help()
			return
		}
		params["name"] = args[0]
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

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "channel_create", params)
}
