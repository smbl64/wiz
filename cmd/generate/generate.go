package generate

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"wiz/internal/generate"
	"wiz/internal/patmgr"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate [flags] <input text>",
	Aliases: []string{"g"},
	Short:   "Call LLM to generate text [alias: g]",
	RunE: func(cmd *cobra.Command, args []string) error {
		var prompt string

		if len(args) > 0 {
			prompt = strings.Join(args, " ")
		} else {
			si, err := io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}

			prompt = strings.TrimSuffix(string(si), "\n")
		}

		temperature, _ := cmd.Flags().GetFloat64("temperature")
		topP, _ := cmd.Flags().GetFloat64("top-p")
		freqPenalty, _ := cmd.Flags().GetFloat64("frequency-penalty")
		presencePenalty, _ := cmd.Flags().GetFloat64("presence-penalty")

		model, _ := cmd.Flags().GetString("model")
		if model == "" {
			model = "qwen2.5:latest" // TODO default model
		}

		gen := generate.Generator{
			Model:            model,
			Temperature:      temperature,
			TopP:             topP,
			FrequencyPenalty: freqPenalty,
			PresencePenalty:  presencePenalty,
		}

		resp, err := gen.Generate(context.Background(), prompt)
		if err != nil {
			return err
		}

		fmt.Println(resp)

		return nil
	},
}

func Initialize(rootCmd *cobra.Command) {
	generateCmd.Flags().String("model", "", "Specify the model to use")
	generateCmd.Flags().Bool("dry-run", false, "Show what would be sent to the model without actually sending it")
	generateCmd.Flags().BoolP("stream", "s", false, "Stream the generated response from the model")
	generateCmd.Flags().StringP("pattern", "p", "", "Pattern to use")
	generateCmd.Flags().Float64P("temperature", "t", 0.7, "Set the temperature")
	generateCmd.Flags().Float64P("top-p", "T", 0.9, "Set the top P")
	generateCmd.Flags().Float64("frequency-penalty", 0.0, "Set the frequencey penalty")
	generateCmd.Flags().Float64("presence-penalty", 0.0, "Set the presence penalty")

	generateCmd.RegisterFlagCompletionFunc("pattern", listPatterns)

	rootCmd.AddCommand(generateCmd)
}

func listPatterns(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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
