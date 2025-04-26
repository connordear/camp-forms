package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"server/middleware"

	"github.com/jackc/pgx/v5"
)

var PORT = ":" + string(os.Getenv("SERVER_PORT"))

type Camp struct {
	Name string `json:"name"`
}

func getCamps(w http.ResponseWriter, r *http.Request) {
	var camps []Camp

	cnn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	rows, cnnErr := cnn.Query(context.Background(), "SELECT name FROM camps")
	if cnnErr != nil {
		log.Fatal(cnnErr)
	}

	camps, mapErr := pgx.CollectRows(rows, pgx.RowToStructByName[Camp])
	if mapErr != nil {
		log.Fatal(mapErr)
	}

	campJson, jsonErr := json.Marshal(camps)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	w.Write(campJson)

	defer rows.Close()
	defer cnn.Close(context.Background())
}

func main() {
	runMigrations()

	router := http.NewServeMux()

	router.HandleFunc("GET /camps", getCamps)

	server := http.Server{
		Addr:    PORT,
		Handler: middleware.Logging(router),
	}

	log.Println("Listening on port", PORT)
	server.ListenAndServe()
}
