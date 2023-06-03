package gopher

import (
	"context"
	"fmt"

	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) AI() *openai.Client {
	return s.ai
}

const openAiModel = openai.GPT3Dot5Turbo

func (s *Service) AIPrompt(ctx context.Context, prompt string, args ...any) (string, error) {
	request := openai.ChatCompletionRequest{
		Model: openAiModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}

	fullPrompt := fmt.Sprintf(prompt, args...)

	request.Messages = append(request.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: fullPrompt,
	})

	// TODO: it seems to fall apart with close to 4000 tokens, the response is sometime just the word "the"
	if err := s.aiTruncateTokens(ctx, request.Messages, 2000); err != nil {
		return "", err
	}

	response, err := s.ai.CreateChatCompletion(context.Background(), request)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}

func (s *Service) aiTruncateTokens(ctx context.Context, messages []openai.ChatCompletionMessage, maxTokens int) error {
	tokenCount, err := s.aiCountTokens(ctx, messages)
	if err != nil {
		return err
	}

	tkm, err := tiktoken.EncodingForModel(openAiModel)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)

		return err
	}

	// Handle truncation of the tokens
	lastMessageIndex := len(messages) - 1
	tokensOver := tokenCount - maxTokens

	if tokensOver > 0 {
		lastMessageTokens := tkm.Encode(messages[lastMessageIndex].Content, nil, nil)
		truncatedTokens := lastMessageTokens[0 : len(lastMessageTokens)-tokensOver-5]
		messages[lastMessageIndex].Content = tkm.Decode(truncatedTokens)
	}

	return nil
}

func (s *Service) aiCountTokens(_ context.Context, messages []openai.ChatCompletionMessage) (int, error) {
	tkm, err := tiktoken.EncodingForModel(openAiModel)
	if err != nil {
		err = fmt.Errorf("EncodingForModel: %v", err)

		return 0, err
	}

	var numTokens int

	var tokensPerMessage int

	var tokensPerName int

	if openAiModel == "gpt-3.5-turbo-0301" || openAiModel == "gpt-3.5-turbo" {
		tokensPerMessage = 4
		tokensPerName = -1
	} else {
		tokensPerMessage = 3
		tokensPerName = 1
	}

	for _, message := range messages {
		numTokens += tokensPerMessage
		numTokens += len(tkm.Encode(message.Content, nil, nil))
		numTokens += len(tkm.Encode(message.Role, nil, nil))
		numTokens += len(tkm.Encode(message.Name, nil, nil))

		if message.Name != "" {
			numTokens += tokensPerName
		}
	}

	numTokens += 3

	return numTokens, nil
}
