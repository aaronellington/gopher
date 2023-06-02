package gopher

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func (s *Service) AI() *openai.Client {
	return s.ai
}

func (s *Service) AIPrompt(ctx context.Context, prompt string, args ...any) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}

	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fmt.Sprintf(prompt, args...),
	})

	resp, err := s.ai.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
