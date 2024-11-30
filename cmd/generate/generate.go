package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/smbl64/wiz/internal/config"
	"github.com/smbl64/wiz/internal/generate"
	"github.com/smbl64/wiz/internal/patmgr"
	"github.com/smbl64/wiz/internal/util/flags"
	"github.com/smbl64/wiz/internal/util/terminal"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "gen [flags] <input text>",
	Aliases: []string{"g"},
	Short:   "Call LLM to generate text [alias: g]",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := strings.Join(args, " ")

		stdinData, err := terminal.ReadStdinIfData()
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		prompt = fmt.Sprintf("%s\n%s", prompt, stdinData)

		temperature, _ := cmd.Flags().GetFloat64("temperature")
		topP, _ := cmd.Flags().GetFloat64("top-p")
		freqPenalty, _ := cmd.Flags().GetFloat64("frequency-penalty")
		presencePenalty, _ := cmd.Flags().GetFloat64("presence-penalty")
		stream, _ := cmd.Flags().GetBool("stream")

		model, _ := cmd.Flags().GetString("model")
		if model == "" {
			model = config.Current().Model
		}

		// Read pattern if required
		var system string
		patName, _ := cmd.Flags().GetString("pattern")
		if patName != "" {
			pattern, err := patmgr.Default().Load(patName)
			if err != nil {
				cmd.PrintErrf("cannot read pattern: %v", err)
				return
			}

			system = pattern
		}

		gen := generate.Generator{
			Model:            model,
			Temperature:      temperature,
			TopP:             topP,
			FrequencyPenalty: freqPenalty,
			PresencePenalty:  presencePenalty,
		}

		if stream {
			gen.StreamingFunc = func(chunk string) error {
				fmt.Fprint(cmd.OutOrStdout(), chunk)
				return nil
			}
		}

		resp, err := gen.Generate(context.Background(), system, prompt)
		if err != nil {
			cmd.PrintErrf("failed to generate content: %v", err)
			return
		}

		if !stream {
			fmt.Fprint(cmd.OutOrStdout(), resp)
		}

		fmt.Fprint(cmd.OutOrStdout(), "\n")
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

	generateCmd.RegisterFlagCompletionFunc("pattern", flags.PatternsFlagCompletionFunc)

	rootCmd.AddCommand(generateCmd)
}
