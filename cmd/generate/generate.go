package generate

import (
	"io"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Call LLM to generate text",
	RunE: func(cmd *cobra.Command, args []string) error {
		var input string

		if len(args) > 0 {
			input = strings.Join(args, " ")
		} else {
			si, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			input = strings.TrimSuffix(string(si), "\n")
		}

		pp.Println(input)
		return nil
	},
}

func Initialize(rootCmd *cobra.Command) {
	rootCmd.AddCommand(generateCmd)
}
