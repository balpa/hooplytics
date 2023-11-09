package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type payload struct {
	Name string `json:"name"`
}

func welcomeGet(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w)

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Welcome!!!")
	fmt.Println("Endpoint Hit: welcomeGet")
}

func welcomePost(w http.ResponseWriter, r *http.Request) {
	setCorsHeaders(w)

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
	http.HandleFunc("/api/welcomePost", welcomePost)
	http.HandleFunc("/api/welcomeGet", welcomeGet)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000/api/welcome")
	w.Header().Set("Access-Control-Max-Age", "10")
}

func main() {
	handleRequests()
}
