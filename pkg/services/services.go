package services

import (
	"image-softcery/pkg/model"
	"image-softcery/pkg/repositories"
)

type Image interface{
	Upload(path string) (int, error)
	Download(id string) (model.Image, error)
}

type Service struct {
	Image
}

func NewService(repos repositories.Image) *Service{
	return &Service{Image: NewImageService(repos)}
}
