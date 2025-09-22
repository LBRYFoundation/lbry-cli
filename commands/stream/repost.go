package commands_stream

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandStreamRepost() *cobra.Command {
	stream_repost := &cobra.Command{
		Use:   "repost",
		Short: "Creates a claim that references an existing stream by its claim id.",
		Run:   HandleCommandStreamRepost,
	}

	stream_repost.Flags().String("name", "", "(str) name of the content (can only consist of a-z A-Z 0-9 and -(dash))")
	stream_repost.Flags().Float64("bid", -1, "(decimal) amount to back the claim")
	stream_repost.Flags().String("claim_id", "", "(str) id of the claim being reposted")
	stream_repost.Flags().Bool("allow_duplicate_name", false, "(bool) create new claim even if one already exists with given name. default: false.")
	stream_repost.Flags().String("title", "", "(str) title of the repost")
	stream_repost.Flags().String("description", "", "(str) description of the repost")
	stream_repost.Flags().StringArray("tags", nil, "(list) add repost tags")
	stream_repost.Flags().String("channel_id", "", "(str) claim id of the publisher channel")
	stream_repost.Flags().String("channel_name", "", "(str) name of the publisher channel")
	stream_repost.Flags().String("channel_account_id", "", "(str) one or more account ids for accounts to look in for channel certificates, defaults to all accounts.")
	stream_repost.Flags().String("account_id", "", "(str) account to use for holding the transaction")
	stream_repost.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")
	stream_repost.Flags().StringArray("funding_account_ids", nil, "(list) ids of accounts to fund this transaction")
	stream_repost.Flags().String("claim_address", "", "(str) address where the claim is sent to, if not specified it will be determined automatically from the account")
	stream_repost.Flags().Bool("preview", false, "(bool) do not broadcast the transaction")
	stream_repost.Flags().Bool("blocking", false, "(bool) wait until transaction is in mempool")

	return stream_repost
}

func HandleCommandStreamRepost(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetFloat64, "bid")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "allow_duplicate_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "title")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "description")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "tags")
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
	if len(args) >= 2 {
		_, exists := params["bid"]
		if exists {
			cmd.Help()
			return
		}
		params["bid"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["claim_id"]
		if exists {
			cmd.Help()
			return
		}
		params["claim_id"] = args[2]
	}
	if len(args) > 3 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("stream_repost", params)
}
