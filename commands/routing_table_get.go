package commands

import "github.com/spf13/cobra"

func CreateCommandRoutingTableGet() *cobra.Command {
	routing_table_get := &cobra.Command{
		Use:   "routing_table_get",
		Short: "Get DHT routing information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return routing_table_get
}
