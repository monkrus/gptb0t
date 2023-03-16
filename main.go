package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	// Set up OpenAI API credentials
	apiKey := "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"
	client := openai.NewClient(apiKey)

	// Prompt the user for input
	var input string
	fmt.Print("You: ")
	fmt.Scanln(&input)

	// Call the OpenAI API to generate a response
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: input,
				},
			},
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 1000,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Print the response from ChatGPT
	var generatedText string
	if len(response.Choices) > 0 {
		if response.Choices[0].Message.Content != "" {
			generatedText = response.Choices[0].Message.Content
		} else {
			generatedText = response.Choices[0].Message.Content
		}
	}
	fmt.Printf("ChatGPT: %v\n", generatedText)
}
