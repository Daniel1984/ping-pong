package models

// Message - holds client to server connection payload structure
type Message struct {
	UserID  int   `json:"user_id"`
	Friends []int `json:"friends"`
}
