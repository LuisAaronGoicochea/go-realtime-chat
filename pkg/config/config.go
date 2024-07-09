package config

import (
    "log"
    "os"
)

var (
    MongoDBURI string
)

func LoadConfig() {
    MongoDBURI = os.Getenv("MONGODB_URI")
    if MongoDBURI == "" {
        log.Fatal("MONGODB_URI not set")
    }
}
