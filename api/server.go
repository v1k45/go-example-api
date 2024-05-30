package api

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/v1k45/shitpost/db"
)

func NewServer(addr string, databaseUrl string) *http.Server {
	conn, err := db.Open(databaseUrl)
	if err != nil {
		slog.Error("database_connection_failed", "error", err)
		log.Fatalf("Error connecting to database: %v", err)
	}

	handler := NewHandler(conn)

	return &http.Server{
		Addr:    addr,
		Handler: handler.Routes(),
	}
}
