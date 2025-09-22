package commands_channel

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandChannelSign() *cobra.Command {
	channel_sign := &cobra.Command{
		Use:   "sign",
		Short: "Signs data using the specified channel signing key.",
		Run:   HandleCommandChannelSign,
	}

	channel_sign.Flags().String("channel_name", "", "(str) name of channel used to sign (or use channel id)")
	channel_sign.Flags().String("channel_id", "", "(str) claim id of channel used to sign (or use channel name)")
	channel_sign.Flags().String("hexdata", "", "(str) data to sign, encoded as hexadecimal")
	channel_sign.Flags().String("salt", "", "(str) salt to use for signing, default is to use timestamp")
	channel_sign.Flags().String("channel_account_id", "", "(str) one or more account ids for accounts to look in for channel certificates, defaults to all accounts.")
	channel_sign.Flags().String("wallet_id", "", "(str) restrict operation to specific wallet")

	return channel_sign
}

func HandleCommandChannelSign(cmd *cobra.Command, args []string) {
	// Create parameter map
	params := map[string]any{}
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_name")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "hexdata")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "salt")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "channel_account_id")
	rpc.AddParameter(params, cmd.Flags(), cmd.Flags().GetString, "wallet_id")

	// Check for arguments
	if len(args) >= 1 {
		_, exists := params["channel_name"]
		if exists {
			cmd.Help()
			return
		}
		params["channel_name"] = args[0]
	}
	if len(args) >= 2 {
		_, exists := params["channel_id"]
		if exists {
			cmd.Help()
			return
		}
		params["channel_id"] = args[1]
	}
	if len(args) >= 3 {
		_, exists := params["hexdata"]
		if exists {
			cmd.Help()
			return
		}
		params["hexdata"] = args[2]
	}
	if len(args) >= 4 {
		_, exists := params["salt"]
		if exists {
			cmd.Help()
			return
		}
		params["salt"] = args[3]
	}
	if len(args) > 4 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("channel_sign", params)
}
