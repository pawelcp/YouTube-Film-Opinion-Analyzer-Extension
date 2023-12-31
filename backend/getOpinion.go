package main

import (
	"encoding/json"
	"fmt"
	"log"

	"os"

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

func getGptOpinion(commentsFormatted commentsFormatted) (string, error) {
	loadEnv()

	apiKey := os.Getenv("OPENAI_API_KEY")
	client := resty.New()

	content := fmt.Sprintf("Based on the provided YouTube video comments, analyze the video's content utility and viewer reception. Rate its likely content quality on a scale of 1-10 and explain your reasoning. Write it in 70 words. '%s", commentsFormatted)

	requestBody := map[string]interface{}{
		"model":      "gpt-3.5-turbo",
		"messages":   []interface{}{map[string]interface{}{"role": "system", "content": content}},
		"max_tokens": 150,
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
	return responseData.Choices[0].Message.Content, nil
}
