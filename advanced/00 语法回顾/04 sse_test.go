package ch00

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

// https://api-docs.deepseek.com/zh-cn/api/create-chat-completion
const (
	BaseURL = "https://api.deepseek.com"
)

type RequestBody struct {
	Messages []map[string]string `json:"messages"`
	Model    string              `json:"model"`
	Stream   bool                `json:"stream"`
}

type ResponseDelta struct {
	Content string `json:"content"`
}

type ResponseChoice struct {
	Delta ResponseDelta `json:"delta"`
}

type ChatCompletion struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Choices []ResponseChoice `json:"choices"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found - using system environment variables")
	}
}

func TestDeepSeekSSE(t *testing.T) {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		t.Fatal("DEEPSEEK_API_KEY environment variable not set")
	}

	url := BaseURL + "/chat/completions"

	reqBody := RequestBody{
		Messages: []map[string]string{
			{"role": "system", "content": "You are a warm supporter"},
			{"role": "user", "content": "我好难受"},
		},
		Model:  "deepseek-chat",
		Stream: true,
	}

	payload, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("API request failed: %d\n%s", resp.StatusCode, body)
	}

	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data:") {
			data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

			if data == "[DONE]" {
				break
			}

			var chunk ChatCompletion
			if err := json.Unmarshal([]byte(data), &chunk); err != nil {
				t.Logf("JSON parse error: %v", err)
				continue
			}

			for _, choice := range chunk.Choices {
				if content := choice.Delta.Content; content != "" {
					fmt.Print(content)
					fullResponse.WriteString(content)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading response: %v", err)
	}

	output := map[string]string{"response": fullResponse.String()}
	file, err := os.Create("output.json")
	require.NoError(t, err)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	require.NoError(t, encoder.Encode(output))

	fmt.Println("\nOutput saved to output.json")
}
