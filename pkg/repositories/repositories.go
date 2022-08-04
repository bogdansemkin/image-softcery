package repositories

import (
	"github.com/jmoiron/sqlx"
	"image-softcery/pkg/model"
)

type Image interface{
	Upload(path string) (int, error)
	Download(id int) (model.Image, error)
}

type Repository struct {
	Image
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{Image: NewImageRepos(db)}
}
