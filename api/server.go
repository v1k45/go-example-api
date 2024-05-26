package api

import "net/http"

func NewServer(addr string) *http.Server {
	handler := NewHandler()

	return &http.Server{
		Addr:    addr,
		Handler: handler.Routes(),
	}
}
