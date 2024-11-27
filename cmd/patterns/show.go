package patterns

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"wiz/internal/patmgr"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:                   "show <pattern>",
	DisableFlagsInUseLine: true,
	Short:                 "Show pattern",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		mgr := patmgr.Default()

		list, err := mgr.List()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		res := lo.Filter(list, func(item string, _ int) bool {
			return strings.HasPrefix(item, toComplete)
		})

		return res, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("No pattern is specified")
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
