package commands_utxo

import "github.com/spf13/cobra"

func CreateCommandUTXORelease() *cobra.Command {
	utxo_release := &cobra.Command{
		Use:   "release",
		Short: "When spending a UTXO it is locally locked to prevent double spends; occasionally this can result in a UTXO being locked which ultimately did not get spent (failed to broadcast, spend transaction was not accepted by blockchain node, etc). This command releases the lock on all UTXOs in your account.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return utxo_release
}
