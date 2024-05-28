package api

import (
	"database/sql"
	"encoding/json"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/v1k45/shitpost/db"
)

const (
	// DefaultPageSize is the default number of shitposts to return per page.
	DefaultPageSize = 5
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

type PaginatedShitpostsResponse struct {
	Results     []db.ListShitpostsRow `json:"results"`
	Count       int                   `json:"count"`
	CurrentPage int                   `json:"currentPage"`
	Pages       int                   `json:"pages"`
}

func (h *Handler) ListShitposts(w http.ResponseWriter, r *http.Request) {
	// get page number
	pageQuery := r.URL.Query().Get("page")
	if pageQuery == "" {
		pageQuery = "1"
	}
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid page number"})
		return
	}

	totalShitposts, err := h.queries.CountShitposts(r.Context())
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Error fetching shitposts"})
		return
	}

	// validate page
	totalPages := int64(math.Ceil(float64(totalShitposts) / float64(DefaultPageSize)))
	if page < 1 || int64(page) > totalPages {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid page number"})
		return
	}

	offset := (page - 1) * DefaultPageSize

	shitposts, err := h.queries.ListShitposts(r.Context(), db.ListShitpostsParams{Limit: int64(DefaultPageSize), Offset: int64(offset)})
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, ErrorResponse{Error: "Error fetching shitposts"})
		return
	}

	if shitposts == nil {
		shitposts = []db.ListShitpostsRow{}
	}

	JSONResponse(w, http.StatusOK, PaginatedShitpostsResponse{
		Results:     shitposts,
		Count:       int(totalShitposts),
		CurrentPage: page,
		Pages:       int(totalPages),
	})
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
