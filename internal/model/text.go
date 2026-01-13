package model

import "time"

type Text struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expires_at"`
}
