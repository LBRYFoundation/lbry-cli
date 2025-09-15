package commands

import "fmt"
import "github.com/spf13/cobra"

func CreateCommandAccountAdd() *cobra.Command {
	account_add := &cobra.Command{
		Use:   "add",
		Short: "Add a previously created account from a seed, private key or public key (read-only). Specify --single_key for single address or vanity address accounts.",
		Run: func(cmd *cobra.Command, args []string) {
			account_name, err := cmd.Flags().GetString("account_name")
			fmt.Println(account_name)
			fmt.Println("====================")
			fmt.Println(err)
		},
	}

	account_add.Flags().String("account_name", "", "(str) name of the account to add")
	account_add.Flags().String("seed", "", "(str) seed to generate new account from")
	account_add.Flags().String("private_key", "", "(str) private key for new account")
	account_add.Flags().String("public_key", "", "(str) public key for new account")
	account_add.Flags().Bool("single_key", false, "(bool) create single key account, default is multi-key")
	account_add.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return account_add
}
