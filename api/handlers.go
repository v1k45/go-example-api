package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/v1k45/shitpost/db"
)

type Handler struct {
	queries *db.Queries
	db      *sql.DB
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

func (h *Handler) ListShitposts(w http.ResponseWriter, r *http.Request) {
	// todo: pagination
	shitposts, err := h.queries.ListShitposts(r.Context())
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Error fetching shitposts"})
		return
	}

	if shitposts == nil {
		shitposts = []db.ListShitpostsRow{}
	}

	JSONResponse(w, http.StatusOK, shitposts)
}

func (h *Handler) CreateShitpost(w http.ResponseWriter, r *http.Request) {
	// Validate the request payload
	var shitpostPayload CreateShitpostPayload
	if err := json.NewDecoder(r.Body).Decode(&shitpostPayload); err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Error decoding shitpost payload"})
		return
	} else if err := shitpostPayload.Validate(); err != nil {
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create the shitpost
	shitpost, err := h.queries.CreateShitpost(r.Context(), db.CreateShitpostParams{
		Author:   shitpostPayload.Author,
		Content:  shitpostPayload.Content,
		Passcode: RandomString(8),
	})
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating shitpost"})
		return
	}

	JSONResponse(w, http.StatusCreated, shitpost)
}

func (h *Handler) GetShitpost(w http.ResponseWriter, r *http.Request) {
	// Extract the shitpost ID from the request path
	postId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid shitpost ID"})
		return
	}

	// Get the shitpost
	shitpost, err := h.queries.GetShitpostById(r.Context(), int64(postId))
	if err != nil {
		JSONResponse(w, http.StatusNotFound, ErrorResponse{Error: "Shitpost not found"})
		return
	}

	JSONResponse(w, http.StatusOK, shitpost)
}

func (h *Handler) DeleteShitpost(w http.ResponseWriter, r *http.Request) {
	// Extract the shitpost ID from the request path
	postId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid shitpost ID"})
		return
	}

	// Validate the request payload
	var deletePayload DeleteShitpostPayload
	if err := json.NewDecoder(r.Body).Decode(&deletePayload); err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	} else if err := deletePayload.Validate(); err != nil {
		JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Verify that the shitpost exists and the passcode is correct
	if _, err = h.queries.GetShitpostByIdAndPasscode(r.Context(), db.GetShitpostByIdAndPasscodeParams{
		ID:       int64(postId),
		Passcode: deletePayload.Passcode,
	}); err != nil {
		JSONResponse(w, http.StatusNotFound, ErrorResponse{Error: "Shitpost id or passcode is incorrect"})
		return
	}

	// Delete the shitpost
	if err = h.queries.DeleteShitpostById(r.Context(), db.DeleteShitpostByIdParams{
		ID:       int64(postId),
		Passcode: deletePayload.Passcode,
	}); err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete shitpost"})
		return
	}

	JSONResponse(w, http.StatusNoContent, nil)
}

// Routes returns the HTTP routes for the API.
func (h *Handler) Routes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("GET /", h.Index)

	routes.HandleFunc("GET /shitposts", h.ListShitposts)
	routes.HandleFunc("POST /shitposts", h.CreateShitpost)
	routes.HandleFunc("GET /shitposts/{id}", h.DeleteShitpost)
	routes.HandleFunc("DELETE /shitposts/{id}", h.DeleteShitpost)
	return routes
}

func NewHandler(conn *sql.DB) *Handler {
	return &Handler{
		db:      conn,
		queries: db.New(conn),
	}
}
