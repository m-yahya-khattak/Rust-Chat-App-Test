package main

import (
	"encoding/json"
	"log"
	"os"
)

// File to store chat messages
const messageFile = "messages.log"

// SaveMessageToFile saves a message to the messages.log file in JSON format
func saveMessageToFile(msg Message) {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(messageFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening message file: %v", err)
		return
	}
	defer file.Close()

	// Marshal the message into JSON
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshalling message: %v", err)
		return
	}

	// Write the JSON message to the file, followed by a newline
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		log.Printf("Error writing to message file: %v", err)
	}
}

// GetMessagesFromFile retrieves all messages for a specific room from the file
func getMessagesFromFile(room string) ([]Message, error) {
	// Open the file in read mode
	file, err := os.Open(messageFile)
	if err != nil {
		log.Printf("Error opening message file: %v", err)
		return nil, err
	}
	defer file.Close()

	var messages []Message
	decoder := json.NewDecoder(file)

	// Decode each line (JSON object) in the file
	for {
		var msg Message
		if err := decoder.Decode(&msg); err != nil {
			if err.Error() == "EOF" {
				break // End of file reached
			}
			log.Printf("Error decoding message: %v", err)
			return nil, err
		}

		// Add the message to the list if it matches the requested room
		if msg.Room == room {
			messages = append(messages, msg)
		}
	}

	return messages, nil
}
