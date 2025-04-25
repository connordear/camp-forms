package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthData struct {
	Timestamp int `json:"timestamp"`
}

var PORT = ":" + string(os.Getenv("SERVER_PORT"))

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthChecking!")
	b, err := json.Marshal(HealthData{Timestamp: 2})
	if err != nil {
		log.Fatalf("Unable to marchal to to %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	runMigrations()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health-check", healthCheck)
	log.Println("Listening on port", PORT)
	http.ListenAndServe(PORT, mux)
}
