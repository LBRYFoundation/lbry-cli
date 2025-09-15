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

	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})

	rootCmd.AddCommand(commands.CreateCommandAccount())

	return rootCmd
}
