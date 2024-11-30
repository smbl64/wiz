package patterns

import (
	"strings"

	"github.com/samber/lo"
	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/spf13/cobra"
)

func patternFlagCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	mgr := patmgr.Default()

	list, err := mgr.List()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	res := lo.Filter(list, func(item string, _ int) bool {
		return strings.HasPrefix(item, toComplete)
	})

	return res, cobra.ShellCompDirectiveNoFileComp
}
