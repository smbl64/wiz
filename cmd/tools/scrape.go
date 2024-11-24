package tools

import "github.com/spf13/cobra"

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape a website and print the text",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
