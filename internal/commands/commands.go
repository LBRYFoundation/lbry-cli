package commands

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"
)

func CreateCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "lbry-cli",
		Short: "An interface to the LBRY Network.",
		Run:   HandleCommand,
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

	rootCmd.AddCommand(CreateCommandAccount())
	rootCmd.AddCommand(CreateCommandAddress())
	rootCmd.AddCommand(CreateCommandBlob())
	rootCmd.AddCommand(CreateCommandChannel())
	rootCmd.AddCommand(CreateCommandClaim())
	rootCmd.AddCommand(CreateCommandCollection())
	rootCmd.AddCommand(CreateCommandFFmpegFind())
	rootCmd.AddCommand(CreateCommandFile())
	rootCmd.AddCommand(CreateCommandGet())
	rootCmd.AddCommand(CreateCommandPeer())
	rootCmd.AddCommand(CreateCommandPreference())
	rootCmd.AddCommand(CreateCommandPublish())
	rootCmd.AddCommand(CreateCommandPurchase())
	rootCmd.AddCommand(CreateCommandResolve())
	rootCmd.AddCommand(CreateCommandRoutingTableGet())
	rootCmd.AddCommand(CreateCommandSettings())
	rootCmd.AddCommand(CreateCommandStart())
	rootCmd.AddCommand(CreateCommandStatus())
	rootCmd.AddCommand(CreateCommandStop())
	rootCmd.AddCommand(CreateCommandStream())
	rootCmd.AddCommand(CreateCommandSupport())
	rootCmd.AddCommand(CreateCommandSync())
	rootCmd.AddCommand(CreateCommandTraceMAlloc())
	rootCmd.AddCommand(CreateCommandTransaction())
	rootCmd.AddCommand(CreateCommandTXO())
	rootCmd.AddCommand(CreateCommandUTXO())
	rootCmd.AddCommand(CreateCommandVersion())
	rootCmd.AddCommand(CreateCommandWallet())

	rootCmd.PersistentFlags().String("api", "", "Host name and port for lbrynet daemon API.")

	rootCmd.Flags().BoolP("version", "V", false, "Show lbrynet CLI version and exit.")

	return rootCmd
}

func HandleCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("version") {
		info, _ := debug.ReadBuildInfo()
		fmt.Println("LBRY CLI " + info.Main.Version)
		return
	}

	cmd.Help()
}
