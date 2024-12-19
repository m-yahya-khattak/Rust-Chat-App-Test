package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func authenticateUser(authData User) bool {
	mutex.Lock()
	defer mutex.Unlock()
	password, exists := users[authData.Username]
	return exists && password == authData.Password
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	if _, exists := users[user.Username]; exists {
		mutex.Unlock()
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	users[user.Username] = user.Password
	mutex.Unlock()

	log.Printf("User %s registered successfully", user.Username)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
