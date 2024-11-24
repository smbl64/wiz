package patterns

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:                   "edit <pattern>",
	Short:                 "Edit a pattern",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
