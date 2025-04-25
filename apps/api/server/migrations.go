package main

import (
	"context"
	"fmt"
	"log"
)

func runMigrations() {
	cnn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	rows, cnnErr := cnn.Query(context.Background(), "SELECT current_version FROM version")

	if cnnErr != nil {
		fmt.Println("Error finding db")
		// run from init ?
	}

	for rows.Next() {
		var current_version string
		err := rows.Scan(&current_version)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Current Database Version is: %s\n", current_version)
	}

	defer rows.Close()
	defer cnn.Close(context.Background())
}
