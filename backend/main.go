package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/senseyeio/roger"
)

type payload struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

func readFile(file string) string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(body)
}

func executeRCode(script string) string {
	rClient, err := roger.NewRClient("127.0.0.1", 6311)
	if err != nil {
		fmt.Println("Failed to create R client: " + err.Error())
		return "Failed to create R client" + err.Error()
	}

	outputRaw, err := rClient.Eval(readFile("./R/for_balpa.R"))
	if err != nil {
		fmt.Println("Command failed: " + err.Error())
		return "Command failed: " + err.Error()
	}

	output, ok := outputRaw.(string)
	if !ok {
		return "Output is not a string"
	}

	return output
}

func executeR(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "Output: %s", executeRCode(payload.Script))
	fmt.Println("Endpoint Hit: executeR")
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
	r.HandleFunc("/api/runRcode", executeR)

	http.ListenAndServe(":3001", nil)
}

func main() {
	handleRequests()
}
