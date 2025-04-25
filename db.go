package main

import (
	"context"
	"fmt"
	"log"
)

func runMigrations() {
	fmt.Println("Running Migrations")
	cnn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	rows, cnnErr := cnn.Query(context.Background(), "SELECT current_version FROM version")

	if cnnErr != nil {
		fmt.Println("Error finding db")
		// run from init ?
	}

	defer rows.Close()
	defer cnn.Close(context.Background())

}
