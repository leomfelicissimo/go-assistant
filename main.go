package main

import (
	"belazap.com/assistant/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load()
	if err != nil {
		log.Error("Unable to load env file")
		return
	}

	server.Initialize()
}
