package repo

import (
	"encoding/json"
	"os"
	"sync"
	"time"

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

func (r *JSONRepo) Save(text model.Text) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[text.ID] = text
	return r.persist()
}

func (r *JSONRepo) Get(id string) (model.Text, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	text, ok := r.store[id]
	if !ok || time.Now().After(text.ExpiresAt) {
		return model.Text{}, false
	}

	return text, true
}

func (r *JSONRepo) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.store, id)
	r.persist()
}

func (r *JSONRepo) load() error {
	file, err := os.Open(r.filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return nil
	}

	defer file.Close()
	return json.NewDecoder(file).Decode(&r.store)
}

func (r *JSONRepo) persist() error {
	tmp := r.filePath + ".tmp"

	f, err := os.Create(tmp)
	if err = json.NewEncoder(f).Encode(r.store); err != nil {
		f.Close()
		return err
	}
	f.Close()
	return os.Rename(tmp, r.filePath)
}

func (r *JSONRepo) cleanup() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		now := time.Now()
		r.mu.Lock()
		for k, v := range r.store {
			if now.After(v.ExpiresAt) {
				delete(r.store, k)
			}
		}
		r.mu.Unlock()
		r.persist()
	}
}
