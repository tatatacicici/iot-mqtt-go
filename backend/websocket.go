package backend

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan interface{})
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebSocketServer(port int) {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Printf("Websocket server started on port %d", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request ke WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		return
	}
	defer func() {
		log.Println("WebSocket connection closed")
		ws.Close()
		delete(clients, ws) // Hapus koneksi dari daftar klien aktif
	}()

	// Tambahkan koneksi baru ke daftar klien
	clients[ws] = true
	log.Printf("New WebSocket connection established. Total clients: %d", len(clients))

	// Loop untuk menjaga koneksi tetap aktif dan membaca pesan
	for {
		_, _, err := ws.ReadMessage() // Membaca pesan dari klien
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket closed unexpectedly: %v", err)
			}
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		message, _ := json.Marshal(msg)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing to client: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Broadcast(data interface{}) {
	broadcast <- data
}
