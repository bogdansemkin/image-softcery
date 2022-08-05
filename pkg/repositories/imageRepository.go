package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"image-softcery/pkg/model"
)

type ImageRepos struct {
	db *sqlx.DB
}

func NewImageRepos(db *sqlx.DB) *ImageRepos{
	return &ImageRepos{db: db}
}

func (r *ImageRepos)Upload(original, resize, halfResize, fullResize string) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s (path,seventy_five_path,half_path,twenty_five_path) VALUES ($1,$2,$3,$4) RETURNING id", imageTable)

	row := r.db.QueryRow(query, original, resize, halfResize, fullResize)

	err := row.Scan(&id)
	if err != nil {
		logrus.Errorf("Error on repo Upload, %s", err)
	}
	return id, nil
}

func (r *ImageRepos) Download(id string) (model.Image, error){
	image := model.Image{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", imageTable)

	err := r.db.Get(&image, query, id)
	if err != nil {
		logrus.Errorf("Error on repo Download, %s", err)
	}
	return image, nil
}
