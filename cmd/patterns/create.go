package patterns

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new pattern",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
