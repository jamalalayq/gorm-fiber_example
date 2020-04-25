package core

// Response generic response.
type Response struct {
	Status string      `json:"status"`
	Code   byte        `json:"code"`
	Err    string       `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
