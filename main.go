package main

import (
	"log"
	backend2 "project-iot/backend"
	backend "project-iot/backend/utils"
)

func main() {
	cfg := backend.LoadConfig()

	backend2.InitMongoDB()

	go backend2.StartWebSocketServer(cfg.WebSocketPort)

	backend2.ConnectAndSubscribe(cfg.MQTTBROKER, cfg.MQTTTOPIC)

	log.Printf("Server running on WebSocket port %d", cfg.WebSocketPort)
	select {}
}
