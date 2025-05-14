package chat

import (
	"fmt"

	"github.com/ak4bento/capsule-shell/internal"
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

var (
	apiURL = internal.GetEnvOrDefault("OPENROUTER_URL", "https://openrouter.ai/api/v1/chat/completions")
	lang   = internal.GetEnvOrDefault("CAPSULE_SHELL_LANGUAGE", "en")
)

func SendMainPrompt(request string) (string, error) {
	apiKey := internal.GetEnvOrDefault("OPENROUTER_API_KEY", "")
	modelRouter := internal.GetEnvOrDefault("MODEL_ROUTER", "meta-llama/llama-4-maverick:free")

	result := map[string]interface{}{}

	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY environment variable is not set")
	}

	prompt := internal.GetMainPrompt(lang)

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
	apiKey := internal.GetEnvOrDefault("OPENROUTER_API_KEY", "")
	result := map[string]interface{}{}

	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	prompt := internal.GetDescriptivePrompt(lang)

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

func SendSatiricalPrompt(request string) (string, error) {
	apiKey := internal.GetEnvOrDefault("OPENROUTER_API_KEY", "")
	result := map[string]interface{}{}

	if apiKey == "" {
		return "", fmt.Errorf("OPENROUTER_API_KEY is not set")
	}

	prompt := internal.GetSatiricalPrompt(lang)

	body := LLMRequest{
		Model: "meta-llama/llama-4-maverick:free", // atau sesuaikan
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
