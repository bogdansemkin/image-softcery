package repositories

import (
	"github.com/jmoiron/sqlx"
	"image-softcery/pkg/model"
)

type Image interface{
	Upload(original, resize, halfResize, fullResize string) (int, error)
	Download(id string) (model.Image, error)
}

type Repository struct {
	Image
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{Image: NewImageRepos(db)}
}
