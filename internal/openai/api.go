package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func MakeAPICall(query, apiKey, model string, temperature float32, templatePath string) string {
	client := NewClient(apiKey)

	var filledQuery string
	var err error

	if templatePath != "" {
		// Load the template from file if specified
		templateStr, err := LoadTemplate(templatePath)
		if err != nil {
			log.Println("Error loading template:", err)
			return "Error: Could not load template"
		}

		// Fill the template with the user query
		filledQuery, err = FillTemplate(templateStr, map[string]string{
			"Query": query,
		})
		if err != nil {
			log.Println("Error filling template:", err)
			return "Error: Could not process template"
		}
	} else {
		// Use the query directly if no template is provided
		filledQuery = query
	}

	response, err := client.DoRequest(filledQuery, model, temperature)
	if err != nil {
		return "Error: " + err.Error()
	}

	return response
}

func MakeAPICallStream(query, apiKey, model string, temperature float32, templatePath string) {
	client := NewClient(apiKey)

	var filledQuery string
	var err error

	if templatePath != "" {
		// Load the template from file if specified
		templateStr, err := LoadTemplate(templatePath)
		if err != nil {
			log.Println("Error loading template:", err)
			fmt.Println("Error: Could not load template")
			return
		}

		// Fill the template with the user query
		filledQuery, err = FillTemplate(templateStr, map[string]string{
			"Query": query,
		})
		if err != nil {
			log.Println("Error filling template:", err)
			fmt.Println("Error: Could not process template")
			return
		}
	} else {
		// Use the query directly if no template is provided
		filledQuery = query
	}

	// Set up the request for streaming
	req := openai.ChatCompletionRequest{
		Model:       model,
		Messages:    []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: filledQuery}},
		Temperature: temperature,
		Stream:      true,
	}

	stream, err := client.Client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
