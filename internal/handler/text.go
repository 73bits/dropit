package handler

import (
	"github.com/73bits/dropit/internal/service"
)

type TextHandler struct {
	service *service.TextService
}

func NewTextHandler(s *service.TextService) *TextHandler {
	return &TextHandler{service: s}
}
