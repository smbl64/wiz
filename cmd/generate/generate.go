package generate

import (
	"io"
	"os"
	"strings"

	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate [flags] <input text>",
	Aliases: []string{"g"},
	Short:   "Call LLM to generate text [alias: g]",
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
	generateCmd.Flags().String("model", "", "Specify the model to use")
	generateCmd.Flags().Bool("dry-run", false, "Show what would be sent to the model without actually sending it")
	generateCmd.Flags().BoolP("stream", "s", false, "Stream the generated response from the model")
	generateCmd.Flags().StringP("pattern", "p", "", "Pattern to use")
	generateCmd.Flags().Float32P("temperature", "t", 0.7, "Set the temperature")
	generateCmd.Flags().Float32P("top-p", "T", 0.9, "Set the top P")
	generateCmd.Flags().Float32("frequency-penalty", 0.0, "Set the frequencey penalty")
	generateCmd.Flags().Float32("presence-penalty", 0.0, "Set the presence penalty")

	rootCmd.AddCommand(generateCmd)
}
