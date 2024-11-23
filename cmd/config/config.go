package config

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configurations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Initialize(rootCmd *cobra.Command) {
	rootCmd.AddCommand(configCmd)
}
