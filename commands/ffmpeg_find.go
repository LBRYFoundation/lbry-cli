package commands

import "github.com/spf13/cobra"

func CreateCommandFFmpegFind() *cobra.Command {
	ffmpeg_find := &cobra.Command{
		Use:   "ffmpeg_find",
		Short: "Get ffmpeg installation information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return ffmpeg_find
}
