package main

import (
	"AssetFlow/database"
	"AssetFlow/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

const shutdownTimeout = 5 * time.Second

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	srv := server.SetupRoutes()

	if err := database.ConnectDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	); err != nil {
		log.Panicf("Failed to initialize and migrate database: %v", err)
	}
	log.Println("migration successful!!")

	go func() {
		if err := srv.Run(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("Failed to start server: %v", err)
		}
	}()

	<-done

	if err := database.CloseDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database closed")

	if err := srv.Shutdown(shutdownTimeout); err != nil {
		log.Panicf("Failed to shutdown server: %v", err)
	}
}
