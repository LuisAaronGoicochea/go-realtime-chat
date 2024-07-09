package main

import (
    "log"
    "net/http"
    "real-time-chat/internal/handlers"
    "real-time-chat/pkg/config"
    "real-time-chat/pkg/database"
)

func main() {
    config.LoadConfig()
    database.ConnectDB()
    http.HandleFunc("/ws", handlers.ChatHandler)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}