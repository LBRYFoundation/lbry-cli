package commands

import "lbry/cli/rpc"
import "github.com/spf13/cobra"

func CreateCommandRoutingTableGet() *cobra.Command {
	routing_table_get := &cobra.Command{
		Use:   "routing_table_get",
		Short: "Get DHT routing information",
		Run:   HandleCommandRoutingTableGet,
	}

	return routing_table_get
}

func HandleCommandRoutingTableGet(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("routing_table_get")
}
