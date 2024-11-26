package patterns

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List patterns",
	Run: func(cmd *cobra.Command, args []string) {
		mgr := makeManager()
		list, err := mgr.List()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %v", err)
			return
		}

		for _, pat := range list {
			fmt.Println(pat)
		}
	},
}
