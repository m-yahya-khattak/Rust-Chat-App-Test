package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	username string
	room     string
}

func removeClient(client *Client) {
	mutex.Lock()
	delete(clients, client)
	mutex.Unlock()
	log.Printf("User %s disconnected", client.username)
}
