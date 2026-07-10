package server

import (
	"AssetFlow/handler"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handler.RegisterUser)

	return mux
}
