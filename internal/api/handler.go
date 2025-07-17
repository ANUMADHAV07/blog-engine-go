package api

import (
	"encoding/json"
	"net/http"

	"github.com/ANUMADHAV07/blog-engine-go.git/internal/blog"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Manager *blog.Manager
}

func NewHandler(manager *blog.Manager) *Handler {
	return &Handler{
		Manager: manager,
	}
}

func (h *Handler) GetPostHandler(w http.ResponseWriter, r *http.Request) {

	slug := chi.URLParam(r, "slug")

	if slug == "" {
		http.Error(w, "slug parameter is required", http.StatusBadRequest)
		return
	}

	post, err := h.Manager.GetPost(slug)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	if slug == "" {
		http.Error(w, "slug parameter is required", http.StatusBadRequest)
		return
	}

	post, err := h.Manager.GetPostByID(slug)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
