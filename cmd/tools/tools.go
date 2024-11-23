package tools

import "github.com/spf13/cobra"

func Initialize(rootCmd *cobra.Command) {
	var toolsCmd = &cobra.Command{
		Use:   "tools",
		Short: "Use built-in tools",
	}
	rootCmd.AddCommand(toolsCmd)

	toolsCmd.AddCommand(youtubeCmd)

}
