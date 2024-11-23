package patterns

import (
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a pattern",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
