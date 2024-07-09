package handlers

import (
    "encoding/json"
    "net/http"
    "real-time-chat/internal/models"
    "real-time-chat/internal/services"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
        return
    }

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }

        var msg models.Message
        if err := json.Unmarshal(p, &msg); err != nil {
            return
        }

        services.SaveMessage(msg)

        if err := conn.WriteMessage(messageType, p); err != nil {
            return
        }
    }
}