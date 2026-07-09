package main

import (
	"AssetFlow/database"
	"fmt"
	"net/http"
)

func main() {
	database.ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Asset Flow is running")
	})
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
