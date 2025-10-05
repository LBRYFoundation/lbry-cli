package commands

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandResolve() *cobra.Command {
	resolve := &cobra.Command{
		Use:   "resolve",
		Short: "Get the claim that a URL refers to.",
		Run:   HandleCommandResolve,
	}

	//resolve.Flags().StringArray("urls", nil, "(str, list) one or more urls to resolve")
	resolve.Flags().String("wallet_id", "", "(str) wallet to check for claim purchase receipts")
	resolve.Flags().String("new_sdk_server", "", "(str) URL of the new SDK server (EXPERIMENTAL)")
	resolve.Flags().Bool("include_purchase_receipt", false, "(bool) lookup and include a receipt if this wallet has purchased the claim being resolved")
	resolve.Flags().Bool("include_is_my_output", false, "(bool) lookup and include a boolean indicating if claim being resolved is yours")
	resolve.Flags().Bool("include_sent_supports", false, "(bool) lookup and sum the total amount of supports you've made to this claim")
	resolve.Flags().Bool("include_sent_tips", false, "(bool) lookup and sum the total amount of tips you've made to this claim (only makes sense when claim is not yours)")
	resolve.Flags().Bool("include_received_tips", false, "(bool) lookup and sum the total amount of tips you've received to this claim (only makes sense when claim is yours)")

	return resolve
}

func HandleCommandResolve(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	//rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetStringArray, "urls")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "new_sdk_server")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_purchase_receipt")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_is_my_output")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_sent_supports")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_sent_tips")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "include_received_tips")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["urls"]
		if exists {
			cmd.Help()
			return
		}
		// Variable arguments
		params["urls"] = args
	}

	// Check for required parameters
	_, exists := params["urls"]
	if !exists {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "resolve", params)
}
