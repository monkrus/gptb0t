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
			Prompt:    input,
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 100,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Print the response from ChatGPT
	fmt.Printf("ChatGPT: %v\n", response.Choices[0].Message.Content)
}
