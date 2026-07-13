package server

import (
	"AssetFlow/handler"
	"AssetFlow/middleware"
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	router http.Handler
}

func protected(h http.HandlerFunc) http.Handler {
	return middleware.Authenticate(h)
}

const (
	readTimeout       = 5 * time.Second
	readHeaderTimeout = 3 * time.Second
	writeTimeout      = 5 * time.Second
)

func setupPublicRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /register", handler.RegisterUser)
	mux.HandleFunc("POST /login", handler.LoginUser)

}

func setupPrivateRoutes(mux *http.ServeMux) {
	mux.Handle("GET /me", protected(handler.GetUser))
}

func SetupRoutes() *Server {
	mux := http.NewServeMux()
	setupPublicRoutes(mux)
	setupPrivateRoutes(mux)

	return &Server{
		router: mux,
	}
}

func (svc *Server) Run(port string) error {

	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return svc.server.ListenAndServe()
}

func (svc *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return svc.server.Shutdown(ctx)
}
