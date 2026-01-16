package repo

import (
	"sync"

	"github.com/73bits/dropit/internal/model"
)

type JSONRepo struct {
	mu       sync.RWMutex
	filePath string
	store    map[string]model.Text
}

func NewJSONRepo(filePath string) (*JSONRepo, error) {
	r := &JSONRepo{
		filePath: filePath,
		store:    make(map[string]model.Text),
	}

	if err := r.load(); err != nil {
		return nil, err
	}

	go r.cleanup()
	return r, nil
}
