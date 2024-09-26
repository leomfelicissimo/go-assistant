package assistant

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tmc/langchaingo/llms/openai"
)

type assistantClient struct {
	ctx context.Context
	llm *openai.LLM
}

type AssistantClient interface {
	Answer(prompt string) (string, error)
}

func New(ctx context.Context) AssistantClient {
	llm, err := openai.New(
		openai.WithAPIVersion(os.Getenv("OPENAI_VERSION")),
		openai.WithAPIType(openai.APITypeAzure),
	)
	if err != nil {
		log.WithError(err).Error("Unabled to initialize llm")
		panic(-1)
	}

	return &assistantClient{
		ctx: ctx,
		llm: llm,
	}
}

func (ai *assistantClient) Answer(prompt string) (string, error) {
	return ai.llm.Call(ai.ctx, prompt)
}
