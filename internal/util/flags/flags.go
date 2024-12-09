package flags

import (
	"context"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/smbl64/wiz/internal/ollama"
	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/spf13/cobra"
)

// PatternsFlagCompletionFunc provides shell completion for patterns.
//
// It can be assigned to `ValidArgsFunction`, or used with `cmd.RegisterFlagCompletionFunc`.
func PatternsFlagCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

func ModelsFlagCompletionFunc(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFn()

	models, err := ollama.Default().ListModels(ctx)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	res := lo.Filter(models, func(item string, _ int) bool {
		return strings.HasPrefix(item, toComplete)
	})

	return res, cobra.ShellCompDirectiveNoFileComp
}
