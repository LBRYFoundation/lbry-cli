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

	//resolve.Flags().StringArray("urls", []string{}, "(str, list) one or more urls to resolve")
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
	urls := []string{}
	//urls, _ := cmd.Flags().GetStringArray("urls")

	wallet_id, _ := cmd.Flags().GetString("wallet_id")
	new_sdk_server, _ := cmd.Flags().GetString("new_sdk_server")
	include_purchase_receipt, _ := cmd.Flags().GetBool("include_purchase_receipt")
	include_is_my_output, _ := cmd.Flags().GetBool("include_is_my_output")
	include_sent_supports, _ := cmd.Flags().GetBool("include_sent_supports")
	include_sent_tips, _ := cmd.Flags().GetBool("include_sent_tips")
	include_received_tips, _ := cmd.Flags().GetBool("include_received_tips")

	// Check for arguments
	if len(args) >= 1 {
		urls = args
	}

	if len(urls) == 0 {
		cmd.Help()
		return
	}

	params := map[string]any{
		"urls": urls,
	}
	if wallet_id != "" {
		params["wallet_id"] = wallet_id
	}
	if new_sdk_server != "" {
		params["new_sdk_server"] = new_sdk_server
	}
	if include_purchase_receipt {
		params["include_purchase_receipt"] = include_purchase_receipt
	}
	if include_is_my_output {
		params["include_is_my_output"] = include_is_my_output
	}
	if include_sent_supports {
		params["include_sent_supports"] = include_sent_supports
	}
	if include_sent_tips {
		params["include_sent_tips"] = include_sent_tips
	}
	if include_received_tips {
		params["include_received_tips"] = include_received_tips
	}

	rpc.ExecuteRPCCommand("resolve", params)
}
