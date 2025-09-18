package commands

import "lbry/cli/rpc"
import "github.com/spf13/cobra"

func CreateCommandVersion() *cobra.Command {
	version := &cobra.Command{
		Use:   "version",
		Short: "Get lbrynet API server version information",
		Run:   HandleCommandVersion,
	}

	return version
}

func HandleCommandVersion(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("version")
}
