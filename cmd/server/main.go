package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codytheroux96/chatterbox/internal"
)

func main() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Failed to decode payload", http.StatusBadRequest)
			return
		}

		if err := internal.RegisterUser(payload.Username, payload.Password); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		token, err := internal.LoginUser(payload.Username, payload.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte(token))
	})

	log.Println("Starting server on :8080!")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
