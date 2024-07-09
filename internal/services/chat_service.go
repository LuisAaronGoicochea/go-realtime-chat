package services

import (
    "context"
    "log"
    "real-time-chat/internal/models"
    "real-time-chat/pkg/database"
)

func SaveMessage(msg models.Message) {
    collection := database.DB.Collection("messages")
    _, err := collection.InsertOne(context.Background(), msg)
    if err != nil {
        log.Printf("Could not save message: %v", err)
    }
}
