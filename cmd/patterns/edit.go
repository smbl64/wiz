package patterns

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a pattern",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
