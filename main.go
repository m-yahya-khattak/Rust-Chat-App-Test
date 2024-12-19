package main

import (
	"encoding/json"
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

func handleChatHistory(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for chat history")

	room := r.URL.Query().Get("room")
	if room == "" {
		http.Error(w, "Room parameter is required", http.StatusBadRequest)
		return
	}

	messages, err := getMessagesFromFile(room)
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		log.Printf("Error fetching messages: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
		log.Printf("Error encoding messages: %v", err)
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/register", handleRegistration)
	http.HandleFunc("/history", handleChatHistory)
	go handleMessages()

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
