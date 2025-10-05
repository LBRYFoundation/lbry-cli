package commands_collection

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandCollectionResolve() *cobra.Command {
	collection_resolve := &cobra.Command{
		Use:   "resolve",
		Short: "Resolve claims in the collection.",
		Run:   HandleCommandCollectionResolve,
	}

	collection_resolve.Flags().String("claim_id", "", "(str) claim id of the collection")
	collection_resolve.Flags().String("url", "", "(str) url of the collection")
	collection_resolve.Flags().String("wallet_id", "", "(str) restrict results to specific wallet")
	collection_resolve.Flags().Int("page", -1, "(int) page to return during paginating")
	collection_resolve.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return collection_resolve
}

func HandleCommandCollectionResolve(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "claim_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "url")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "collection_resolve", params)
}
