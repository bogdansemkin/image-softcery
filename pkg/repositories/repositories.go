package repositories

import (
	"github.com/bogdansemkin/image-softcery/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Image interface {
	Upload(original, resize, halfResize, fullResize string) (int, error)
	Download(id string) (model.Image, error)
}

type Repository struct {
	Image
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Image: NewImageRepos(db)}
}
