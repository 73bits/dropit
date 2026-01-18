package handler

import (
	"encoding/json"
	"net/http"

	"github.com/73bits/dropit/internal/service"
)

type TextHandler struct {
	service *service.TextService
}

func NewTextHandler(s *service.TextService) *TextHandler {
	return &TextHandler{service: s}
}

func (h *TextHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Content    string `json:"content"`
		TTLSeconds int    `json:"ttl_seconds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	text, err := h.service.Create(req.Content, req.TTLSeconds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"url": "http://" + r.Host + "/text/" + text.ID,
	})
}
