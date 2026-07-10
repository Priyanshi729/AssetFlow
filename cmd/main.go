package main

import (
	"AssetFlow/database"
	"log"
	"net/http"
)

func main() {

	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	// Close DB when server stops
	defer database.CloseDB()

	// Register routes
	http.HandleFunc("/", HomeHandler)

	log.Println("Server started on :8080")

	// Start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AssetFlow Server Running"))
}
