package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	clients   = make(map[*Client]bool)
	rooms     = make(map[string][]*Client)
	users     = make(map[string]string) // Username -> Password
	broadcast = make(chan Message)
	mutex     = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/register", handleRegistration)
	go handleMessages()

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
