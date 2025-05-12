package chat

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-resty/resty/v2"
)

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type LLMRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

var apiURL = os.Getenv("OPENROUTER_URL")
// var apiURL = "https://openrouter.ai/api/v1/chat/completions"

func SendPrompt(request string) (string, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	modelRouter := os.Getenv("MODEL_ROUTER")

	result := map[string]interface{}{}

	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY environment variable is not set")
	}

	prompt := "You are a helpful assistant in OS " + runtime.GOOS + ". \n" +
		"Your name is Capsule Shell. You are a shell command line interpreter. \n" +
		"You can only respond with shell commands. Do not add any explanation or additional text. \n" +
		"If the user asks for help, just say 'I am capsule shell command line interpreter'."

	body := LLMRequest{
		Model: modelRouter,
		Messages: []Message{
			{
				Role:    "system",
				Content: []Content{{Type: "text", Text: prompt}},
			},
			{
				Role: "user",
				Content: []Content{
					{Type: "text", Text: request},
				},
			},
		},
	}

	client := resty.New()
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("HTTP-Referer", "https://github.com/ak4bento/capsule-shell").
		SetBody(body).
		SetResult(&result).
		Post(apiURL)

	if err != nil {
		return "", err
	}

	if errResp, ok := result["error"]; ok {
		return "", fmt.Errorf("API error: %v", errResp)
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("None choices found in response")
	}

	msg, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("message not found in response")
	}

	content, ok := msg["content"].(string)
	if !ok {
		return "", fmt.Errorf("content not found in response")
	}

	return content, nil
}

func SendDescriptivePrompt(request string) (string, error) {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	result := map[string]interface{}{}

	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	prompt := "You are a helpful AI running in OS " + runtime.GOOS + ".\n" +
		"Your name is Capsule Shell. You are a CLI assistant that explains how to run shell commands.\n" +
		"Respond with clear step-by-step explanation followed by the final command.\n" +
		"Format output like this:\n" +
		"1. Step one\n 2. Step two\n\nShell Command:\n```\n<command here>\n```"

	body := LLMRequest{
		Model: "meta-llama/llama-4-maverick:free",
		Messages: []Message{
			{
				Role:    "system",
				Content: []Content{{Type: "text", Text: prompt}},
			},
			{
				Role: "user",
				Content: []Content{
					{Type: "text", Text: request},
				},
			},
		},
	}

	client := resty.New()
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetHeader("HTTP-Referer", "https://github.com/ak4bento/capsule-shell").
		SetBody(body).
		SetResult(&result).
		Post(apiURL)

	if err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("None choices found in response")
	}

	msg, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Failed parse response message")
	}

	content, ok := msg["content"].(string)
	if !ok {
		return "", fmt.Errorf("Content not found in response")
	}

	return content, nil
}

