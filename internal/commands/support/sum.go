package commands_support

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandSupportSum() *cobra.Command {
	support_sum := &cobra.Command{
		Use:   "sum",
		Short: "List total staked supports for a claim, grouped by the channel that signed the support.",
		Run:   HandleCommandSupportSum,
	}

	support_sum.Flags().String("claim_id", "", "(str)  claim id")
	support_sum.Flags().String("new_sdk_server", "", "(str)  URL of the new SDK server (EXPERIMENTAL)")
	support_sum.Flags().Bool("include_channel_content", false, "(bool) if claim_id is for a channel, include supports for claims in that channel")
	support_sum.Flags().Int("page", -1, "(int)  page to return during paginating")
	support_sum.Flags().Int("page_size", -1, "(int)  number of items on page during pagination")

	return support_sum
}

func HandleCommandSupportSum(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "new_sdk_server")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_channel_content")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

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
		_, exists := params["new_sdk_server"]
		if exists {
			cmd.Help()
			return
		}
		params["new_sdk_server"] = args[1]
	}
	if len(args) > 2 {
		cmd.Help()
		return
	}

	// Check for required parameters
	_, existsClaimID := params["claim_id"]
	if !existsClaimID {
		cmd.Help()
		return
	}
	_, existsNewSDKServer := params["new_sdk_server"]
	if !existsNewSDKServer {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "support_sum", params)
}
