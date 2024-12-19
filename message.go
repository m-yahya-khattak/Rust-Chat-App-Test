package main

import "log"

type Message struct {
	From    string `json:"from"`
	To      string `json:"to,omitempty"`
	Room    string `json:"room,omitempty"`
	Content string `json:"content,omitempty"`
	Action  string `json:"action"` // "join", "leave", "message"
}

func handleMessages() {
	for {
		msg := <-broadcast
		log.Printf("Processing message: %+v", msg)

		mutex.Lock()
		switch msg.Action {
		case "join":
			log.Printf("User %s joining room %s", msg.From, msg.Room)
			for client := range clients {
				if client.username == msg.From {
					rooms[msg.Room] = append(rooms[msg.Room], client)
					log.Printf("User %s added to room %s", msg.From, msg.Room)
				}
			}
		case "leave":
			log.Printf("User %s leaving room %s", msg.From, msg.Room)
			for i, client := range rooms[msg.Room] {
				if client.username == msg.From {
					rooms[msg.Room] = append(rooms[msg.Room][:i], rooms[msg.Room][i+1:]...)
					log.Printf("User %s removed from room %s", msg.From, msg.Room)
					break
				}
			}
		case "message":
			// Save the message to the file
			saveMessageToFile(msg)

			if msg.To != "" {
				// Direct Message (DM)
				log.Printf("Direct message from %s to %s: %s", msg.From, msg.To, msg.Content)
				for client := range clients {
					if client.username == msg.To {
						log.Printf("Sending DM to %s from %s", msg.To, msg.From)
						if err := client.conn.WriteJSON(msg); err != nil {
							log.Printf("Error sending DM to user %s: %v", msg.To, err)
						}
						break
					}
				}
			} else {
				// Broadcast Message
				log.Printf("Broadcasting message to room %s from user %s", msg.Room, msg.From)
				for _, client := range rooms[msg.Room] {
					if client != nil && client.username != msg.From {
						if err := client.conn.WriteJSON(msg); err != nil {
							log.Printf("Error sending message to user %s in room %s: %v", client.username, msg.Room, err)
						}
					}
				}
			}
		}
		mutex.Unlock()
	}
}
