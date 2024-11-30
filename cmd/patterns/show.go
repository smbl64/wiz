package patterns

import (
	"fmt"
	"os"

	"github.com/smbl64/wiz/internal/patmgr"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:                   "show <pattern>",
	Short:                 "Show pattern",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     patternFlagCompletion,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return cmd.Usage()
		}

		patName := args[0]
		mgr := patmgr.Default()

		tldr, err := mgr.Load(patName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to show pattern: %v", err)
			return nil
		}
		fmt.Println(tldr)

		return nil
	},
}
