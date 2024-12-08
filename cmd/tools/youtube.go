package tools

import "github.com/spf13/cobra"

var youtubeCmd = &cobra.Command{
	Use:               "youtube",
	Aliases:           []string{"yt"},
	Short:             "Download video transcript from YouTube",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
