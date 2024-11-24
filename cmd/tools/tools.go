package tools

import "github.com/spf13/cobra"

func Initialize(rootCmd *cobra.Command) {
	var toolsCmd = &cobra.Command{
		Use:     "tools",
		Aliases: []string{"t"},
		Short:   "Use built-in tools [alias: t]",
	}
	rootCmd.AddCommand(toolsCmd)

	toolsCmd.AddCommand(youtubeCmd)
	toolsCmd.AddCommand(scrapeCmd)

}
