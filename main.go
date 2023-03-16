package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Set up OpenAI API credentials
	apiKey := "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"
	client := openai.NewClient(apiKey)

	// Define the chatbot prompt
	prompt := `
As an advanced chatbot, your primary goal is to assist users to the best of your ability. This may involve answering questions, providing helpful information, or completing tasks based on user input. In order to effectively assist users, it is important to be detailed and thorough in your responses. Use examples and evidence to support your points and justify your recommendations or solutions.

%s
User: %s
Chatbot:`

	// Initialize the conversation history
	conversationHistory := ""

	// Start the conversation loop
	for {
		// Prompt the user for input
		fmt.Print("You: ")
		var userInput string
		fmt.Scanln(&userInput)

		// If the user types "exit", break out of the conversation loop
		if userInput == "exit" {
			break
		}

		// Call the OpenAI API to generate a response
		response, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
			Prompt:      fmt.Sprintf(prompt, conversationHistory, userInput),
			MaxTokens:   100,
			Model:       openai.GPT3TextDavinci003,
			Temperature: 0.9,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		// Extract the response text from the response object
		var generatedText string
		if len(response.Choices) > 0 {
			generatedText = response.Choices[0].Text
		}

		// Print the response from the chatbot
		fmt.Printf("Chatbot: %v\n", generatedText)

		// Add the conversation to the conversation history
		conversationHistory += fmt.Sprintf("You: %s\nChatbot: %s\n", userInput, generatedText)
	}
}
