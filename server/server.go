package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

const (
	readTimeout       = 5 * time.Second
	readHeaderTimeout = 3 * time.Second
	writeTimeout      = 5 * time.Second
)

func (svc *Server) Run(port string) error {

	mux := SetupRoutes()

	svc.server = &http.Server{
		Addr:              port,
		Handler:           mux,
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
