package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/go-api/model"
)

func SendQuestion(question string) {
	// API endpoint URL
	//apiUrl := os.Getenv("OPENAI_API_URL")
	//apiKey := os.Getenv("API_KEY")
	apiKey := "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"

	// Request payload
	requestBody := model.RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []model.T{
			{
				Role:    "user",
				Content: question,
			},
		},
		Temperature: 0.7,
		Max_tokens:  100,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers if needed
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	choices := result["choices"].([]interface{})
	firstChoice := choices[0]
	message := firstChoice.(map[string]interface{})["message"]
	content := message.(map[string]interface{})["content"]
	fmt.Println(content)
}
