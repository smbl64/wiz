package generate

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

type Generator struct {
	Model            string
	Temperature      float64
	TopP             float64
	FrequencyPenalty float64
	PresencePenalty  float64
	StreamingFunc    func(chunk string) error
}

func (g *Generator) Generate(ctx context.Context, prompt string) (string, error) {
	llm, err := makeOllamaModel(g.Model)
	if err != nil {
		return "", fmt.Errorf("failed to create ollama client: %v", err)
	}

	m := llms.TextParts(llms.ChatMessageTypeHuman, prompt)

	stfunc := noopStreamingFunc

	// A proxy over the given streaming func, to simplify things
	if g.StreamingFunc != nil {
		stfunc = func(_ context.Context, chunk []byte) error {
			return g.StreamingFunc(string(chunk))
		}
	}

	resp, err := llm.GenerateContent(
		ctx,
		[]llms.MessageContent{m},
		llms.WithTemperature(g.Temperature),
		llms.WithTopP(g.TopP),
		llms.WithFrequencyPenalty(g.FrequencyPenalty),
		llms.WithPresencePenalty(g.PresencePenalty),
		llms.WithStreamingFunc(stfunc),
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Content, nil
}

func noopStreamingFunc(ctx context.Context, chunk []byte) error {
	return nil
}

func makeOllamaModel(modelName string) (llms.Model, error) {
	return ollama.New(ollama.WithModel(modelName))
}

func makeOpenAIModel(token, modelName string) (llms.Model, error) {
	return openai.New(openai.WithToken(token), openai.WithModel(modelName))
}
