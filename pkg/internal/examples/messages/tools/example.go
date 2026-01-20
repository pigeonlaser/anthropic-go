package main

import (
	"context"

	"github.com/pigeonlaser/anthropic-go/v3/pkg/anthropic"
	"github.com/pigeonlaser/anthropic-go/v3/pkg/anthropic/client/native"
)

type WeatherRequest struct {
	City string `json:"city"`
	Unit string `json:"unit"`
}

func main() {
	ctx := context.Background()
	client, err := native.MakeClient(native.Config{
		APIKey: "your-api-key",
	})
	if err != nil {
		panic(err)
	}

	// Prepare a message request
	request := &anthropic.MessageRequest{
		Model:             anthropic.Claude3Opus,
		MaxTokensToSample: 1024,
		Tools: []anthropic.Tool{
			{
				Name:        "get_weather",
				Description: "Get the weather",
			},
		},
		Messages: []anthropic.MessagePartRequest{
			{
				Role: "user",
				Content: []anthropic.ContentBlock{
					anthropic.NewTextContentBlock("what's the weather in Charleston?"),
				},
			},
		},
	}

	// Call the Message method
	response, err := client.Message(ctx, request)
	if err != nil {
		panic(err)
	}

	if response.StopReason == "tool_use" {
		// Do something with the tool response
	}
}
