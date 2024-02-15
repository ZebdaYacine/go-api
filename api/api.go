package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	config "example.com/go-api/config"
	"example.com/go-api/model"
)

func SendQuestion(question string) {
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
	req, err := http.NewRequest("POST", config.Api_url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers if needed
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Api_key)

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
