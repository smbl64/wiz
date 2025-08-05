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
	Use:               "generate [flags] <input text>",
	Aliases:           []string{"g", "gen"},
	Short:             "Call LLM to generate text [alias: g, gen]",
	ValidArgsFunction: cobra.NoFileCompletions,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := strings.TrimSpace(strings.Join(args, " "))

		stdinData, err := terminal.ReadStdinIfData()
		if err != nil {
			cmd.PrintErr(err)
			return
		} else if len(stdinData) > 0 {
			prompt = fmt.Sprintf("%s\n%s", stdinData, prompt)
		}

		prompt = strings.TrimSpace(prompt)
		if len(prompt) == 0 {
			cmd.PrintErr("No message is provided")
			return
		}

		temperature, _ := cmd.Flags().GetFloat64("temperature")
		topP, _ := cmd.Flags().GetFloat64("top-p")
		freqPenalty, _ := cmd.Flags().GetFloat64("frequency-penalty")
		presencePenalty, _ := cmd.Flags().GetFloat64("presence-penalty")
		noStream, _ := cmd.Flags().GetBool("no-stream")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		stream := !noStream

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

		if dryRun {
			cmd.Printf("System:\n%s\n", system)
			cmd.Println()
			cmd.Printf("User:\n%s\n", prompt)
			cmd.Println()
			cmd.Printf("Model: %s\n", model)
			cmd.Println()
			return
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
	generateCmd.Flags().Bool("no-stream", false, "Do not stream the generated response from the model")
	generateCmd.Flags().StringP("pattern", "p", "", "Pattern to use")
	generateCmd.Flags().Float64P("temperature", "t", 0.7, "Set the temperature")
	generateCmd.Flags().Float64P("top-p", "T", 0.9, "Set the top P")
	generateCmd.Flags().Float64("frequency-penalty", 0.0, "Set the frequency penalty")
	generateCmd.Flags().Float64("presence-penalty", 0.0, "Set the presence penalty")

	generateCmd.RegisterFlagCompletionFunc("pattern", flags.PatternsFlagCompletionFunc)
	generateCmd.RegisterFlagCompletionFunc("model", flags.ModelsFlagCompletionFunc)

	rootCmd.AddCommand(generateCmd)
}
