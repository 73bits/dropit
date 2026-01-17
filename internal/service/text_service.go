package service

import (
	"errors"
	"time"

	"github.com/73bits/dropit/internal/model"
	"github.com/73bits/dropit/internal/repo"
	"github.com/73bits/dropit/internal/util"
)

type TextService struct {
	repo repo.TextRepo
}

func NewTextService(r repo.TextRepo) *TextService {
	return &TextService{repo: r}
}

func (s *TextService) Create(content string, ttlseconds int) (model.Text, error) {
	if content == "" || ttlseconds <= 0 {
		return model.Text{}, errors.New("invalid input")
	}

	text := model.Text{
		ID:        util.GenerateID(),
		Content:   content,
		ExpiresAt: time.Now().Add(time.Duration(ttlseconds) * time.Second),
	}

	err := s.repo.Save(text)
	return text, err
}
