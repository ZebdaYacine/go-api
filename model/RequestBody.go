package model

type RequestBody struct {
	Model       string  `json:"model"`
	Messages    []T     `json:"messages"`
	Temperature float32 `json:"temperature"`
	Max_tokens  int64   `json:"max_tokens"`
}

type T struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
