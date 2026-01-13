package repo

import "github.com/73bits/dropit/internal/model"

type TextRepo interface {
	Save(text model.Text) error
	Get(id string) (model.Text, bool)
	Delete(id string)
}
