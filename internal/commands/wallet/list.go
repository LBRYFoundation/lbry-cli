package commands_wallet

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandWalletList() *cobra.Command {
	wallet_list := &cobra.Command{
		Use:   "list",
		Short: "List wallets.",
		Run:   HandleCommandWalletList,
	}

	wallet_list.Flags().String("wallet_id", "", "(str) show specific wallet only")
	wallet_list.Flags().Int("page", -1, "(int) page to return during paginating")
	wallet_list.Flags().Int("page_size", -1, "(int) number of items on page during pagination")

	return wallet_list
}

func HandleCommandWalletList(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetInt, "page_size")

	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "wallet_list", params)
}
