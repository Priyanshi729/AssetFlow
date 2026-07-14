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

func protectedWithRoles(h http.HandlerFunc, roles ...string) http.Handler {
	return middleware.Authenticate(
		middleware.RequireRoles(roles...)(h),
	)
}

const (
	readTimeout       = 5 * time.Second
	readHeaderTimeout = 3 * time.Second
	writeTimeout      = 5 * time.Second
)

func setupPublicRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /v1/register", handler.RegisterUser)
	mux.HandleFunc("POST /v1/login", handler.LoginUser)

}

func setupPrivateRoutes(mux *http.ServeMux) {
	mux.Handle("GET /v1/user/me", protectedWithRoles(handler.GetUser, "admin", "project-manager", "employee"))
	mux.Handle("POST /v1/user/logout", protectedWithRoles(handler.LogoutUser, "admin", "project-manager", "employee"))
	mux.Handle("DELETE /v1/user/{userID}", protectedWithRoles(handler.DeleteUser, "admin", "project-manager", "employee"))

	mux.Handle("POST /v1/assets", protectedWithRoles(handler.CreateAsset, "admin", "project-manager"))
	mux.Handle("GET /v1/assets", protectedWithRoles(handler.GetAssets, "admin", "project-manager"))
	mux.Handle("GET /v1/assets/{assetID}", protectedWithRoles(handler.GetAssetByID, "admin", "project-manager"))

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
