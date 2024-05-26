package api

import (
	"net/http"
	"os"
	"time"
)

type Handler struct {
}

type WelcomeResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Env     string    `json:"env"`
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	response := WelcomeResponse{
		Message: "Hello, World!",
		Time:    time.Now(),
		Env:     hostname,
	}

	JSONResponse(w, http.StatusOK, response)
}

func (h *Handler) Routes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("/", h.Index)
	return routes
}

func NewHandler() *Handler {
	return &Handler{}
}
