package commands

import (
	"lbry/cli/rpc"

	"github.com/spf13/cobra"
)

func CreateCommandFFmpegFind() *cobra.Command {
	ffmpeg_find := &cobra.Command{
		Use:   "ffmpeg_find",
		Short: "Get ffmpeg installation information",
		Run:   HandleCommandFFmpegFind,
	}

	return ffmpeg_find
}

func HandleCommandFFmpegFind(cmd *cobra.Command, args []string) {
	// Check for arguments
	if len(args) > 0 {
		cmd.Help()
		return
	}

	rpc.ExecuteRPCCommand("ffmpeg_find")
}
