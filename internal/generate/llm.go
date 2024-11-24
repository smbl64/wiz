package generate

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

type GenerateRequest struct {
	Prompt          string
	Model           string
	Temperature     float64
	TopP            float64
	FrequencyPenalty     float64
	PresencePenalty float64
}

func Generate(ctx context.Context, req GenerateRequest) (string, error) {
	llm, err := makeOllamaModel(req.Model)
	if err != nil {
		return "", fmt.Errorf("failed to create ollama client: %v", err)
	}

	m := llms.TextParts(llms.ChatMessageTypeHuman, req.Prompt)

	resp, err := llm.GenerateContent(
		ctx,
		[]llms.MessageContent{m},
		llms.WithTemperature(req.Temperature),
		llms.WithTopP(req.TopP),
		llms.WithFrequencyPenalty(req.FrequencyPenalty),
		llms.WithPresencePenalty(req.PresencePenalty),
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Content, nil
}

func makeOllamaModel(modelName string) (llms.Model, error) {
	return ollama.New(ollama.WithModel(modelName))
}

func makeOpenAIModel(token, modelName string) (llms.Model, error) {
	return openai.New(openai.WithToken(token), openai.WithModel(modelName))
}
