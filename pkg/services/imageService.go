package services

import (
	"image-softcery/pkg/model"
	"image-softcery/pkg/repositories"
)

type ImageService struct{
	repo repositories.Image
}

func NewImageService(repos repositories.Image) *ImageService{
	return &ImageService{repo: repos}
}

func (s *ImageService)Upload(path string) (int, error){
	return s.repo.Upload(path)
}

func (s *ImageService) Download(id string) (model.Image, error){
	return s.repo.Download(id)
}