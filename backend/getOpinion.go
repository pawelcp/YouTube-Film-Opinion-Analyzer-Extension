package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type ResponseData struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Model   string `json:"model"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type commentsFormatted []string

func getGptOpinion(commentsFormatted commentsFormatted) {
	apiKey := "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"
	client := resty.New()

	content := fmt.Sprintf("tell me how more or less people have received the movie '%s' based on these comments", commentsFormatted)

	requestBody := map[string]interface{}{
		"model":      "gpt-3.5-turbo",
		"messages":   []interface{}{map[string]interface{}{"role": "system", "content": content}},
		"max_tokens": 50,
	}

	response, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Failed to send the request: %v", err)
	}

	var responseData ResponseData
	err = json.Unmarshal(response.Body(), &responseData)
	if err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	fmt.Println("Assistant's response:", responseData.Choices[0].Message.Content)
}
