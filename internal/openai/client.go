package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type APIClient struct {
	Client *openai.Client
}

func NewClient(apiKey string) *APIClient {
	return &APIClient{
		Client: openai.NewClient(apiKey),
	}
}

func (c *APIClient) DoRequest(prompt string, model string, temperature float32) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Temperature: temperature,
	}

	resp, err := c.Client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", nil
}
