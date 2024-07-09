package models

type Message struct {
    Username  string `json:"username"`
    Content   string `json:"content"`
    Timestamp int64  `json:"timestamp"`
}
