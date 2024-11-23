package patterns

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List patterns",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
