package service

import (
	"github.com/73bits/dropit/internal/repo"
)

type TextService struct {
	repo repo.TextRepo
}

func NewTextService(r repo.TextRepo) *TextService {
	return &TextService{repo: r}
}
