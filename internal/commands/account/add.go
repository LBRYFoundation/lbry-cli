package commands_account

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandAccountAdd() *cobra.Command {
	account_add := &cobra.Command{
		Use:   "add",
		Short: "Add a previously created account from a seed, private key or public key (read-only). Specify --single_key for single address or vanity address accounts.",
		Run:   HandleCommandAccountAdd,
	}

	account_add.Flags().String("account_name", "", "(str) name of the account to add")
	account_add.Flags().String("seed", "", "(str) seed to generate new account from")
	account_add.Flags().String("private_key", "", "(str) private key for new account")
	account_add.Flags().String("public_key", "", "(str) public key for new account")
	account_add.Flags().Bool("single_key", false, "(bool) create single key account, default is multi-key")
	account_add.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return account_add
}

func HandleCommandAccountAdd(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "account_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "seed")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "private_key")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "public_key")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetBool, "single_key")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["account_name"]
		if exists {
			cmd.Help()
			return
		}
		params["account_name"] = args[0]
	}
	if len(args) > 1 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "account_add", params)
}
