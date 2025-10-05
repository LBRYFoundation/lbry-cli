package commands

import (
	"lbry/cli/internal/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandRoutingTableGet() *cobra.Command {
	routing_table_get := &cobra.Command{
		Use:   "routing_table_get",
		Short: "Get DHT routing information",
		Run:   HandleCommandRoutingTableGet,
	}

	return routing_table_get
}

func HandleCommandRoutingTableGet(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand(rpc.GetURL(cmd), "routing_table_get")
}
