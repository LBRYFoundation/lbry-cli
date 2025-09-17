package main

import "fmt"
import "os"
import "lbry/cli/commands"
import "github.com/spf13/cobra"
import cc "github.com/ivanpirog/coloredcobra"

func main() {
	cmd := createCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "lbry-cli",
		Short: "An interface to the LBRY Network.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})

	rootCmd.AddCommand(commands.CreateCommandAccount())
	rootCmd.AddCommand(commands.CreateCommandAddress())
	rootCmd.AddCommand(commands.CreateCommandBlob())
	rootCmd.AddCommand(commands.CreateCommandChannel())
	rootCmd.AddCommand(commands.CreateCommandClaim())
	rootCmd.AddCommand(commands.CreateCommandCollection())
	rootCmd.AddCommand(commands.CreateCommandFFmpegFind())
	rootCmd.AddCommand(commands.CreateCommandFile())
	rootCmd.AddCommand(commands.CreateCommandGet())
	rootCmd.AddCommand(commands.CreateCommandPeer())
	rootCmd.AddCommand(commands.CreateCommandPreference())
	rootCmd.AddCommand(commands.CreateCommandPublish())
	rootCmd.AddCommand(commands.CreateCommandPurchase())
	rootCmd.AddCommand(commands.CreateCommandResolve())
	rootCmd.AddCommand(commands.CreateCommandRoutingTableGet())
	rootCmd.AddCommand(commands.CreateCommandSettings())
	rootCmd.AddCommand(commands.CreateCommandStart())
	rootCmd.AddCommand(commands.CreateCommandStatus())
	rootCmd.AddCommand(commands.CreateCommandStop())
	rootCmd.AddCommand(commands.CreateCommandStream())
	rootCmd.AddCommand(commands.CreateCommandSupport())
	rootCmd.AddCommand(commands.CreateCommandSync())
	rootCmd.AddCommand(commands.CreateCommandTraceMAlloc())
	rootCmd.AddCommand(commands.CreateCommandTransaction())
	rootCmd.AddCommand(commands.CreateCommandTXO())
	rootCmd.AddCommand(commands.CreateCommandUTXO())
	rootCmd.AddCommand(commands.CreateCommandVersion())
	rootCmd.AddCommand(commands.CreateCommandWallet())

	return rootCmd
}
