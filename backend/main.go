package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type payload struct {
	Name string `json:"name"`
}

func welcomeGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Welcome!!!")
	fmt.Println("Endpoint Hit: welcomeGet")
}

func welcomePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload payload

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Welcome to the server, %s!", payload.Name)
	fmt.Println("Endpoint Hit: welcome")
}

func handleRequests() {
	r := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.Handle("/", corsHandler.Handler(r))

	r.HandleFunc("/api/welcomeGet", welcomeGet)
	r.HandleFunc("/api/welcomePost", welcomePost)

	http.ListenAndServe(":3001", nil)
}

func main() {
	handleRequests()
}
