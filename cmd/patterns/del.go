package patterns

import (
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:                   "del <pattern>",
	Short:                 "Delete a pattern",
	Aliases:               []string{"delete"},
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
