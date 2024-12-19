package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	var authData User
	if err := conn.ReadJSON(&authData); err != nil {
		log.Println("Error reading authentication data:", err)
		conn.WriteJSON(map[string]string{"error": "Invalid authentication data"})
		return
	}

	if !authenticateUser(authData) {
		log.Printf("Authentication failed for user %s", authData.Username)
		conn.WriteJSON(map[string]string{"error": "Authentication failed"})
		return
	}

	client := &Client{
		conn:     conn,
		username: authData.Username,
	}

	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	log.Printf("User %s connected", client.username)

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Connection closed for user %s: %v", client.username, err)
			removeClient(client)
			return
		}
		broadcast <- msg
	}
}
