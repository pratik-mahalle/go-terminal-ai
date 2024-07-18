package ai

import (
    "context"
    "fmt"
    "log"
    "os"

    openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func init() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    client = openai.NewClient(apiKey)
}

func GetSuggestions(prompt string) string {
    req := openai.CompletionRequest{
        Model:     openai.GPT3,
        Prompt:    prompt,
        MaxTokens: 100,
    }
    resp, err := client.CreateCompletion(context.Background(), req)
    if err != nil {
        log.Printf("Error getting AI suggestions: %v", err)
        return ""
    }
    return resp.Choices[0].Text
}

