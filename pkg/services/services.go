package services

import (
	"github.com/bogdansemkin/image-softcery/pkg/model"
	"github.com/bogdansemkin/image-softcery/pkg/repositories"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go

type Image interface {
	Upload(original, resize, halfResize, fullResize string) (int, error)
	Download(id string) (model.Image, error)
}

type Service struct {
	Image
}

func NewService(repos repositories.Image) *Service {
	return &Service{Image: NewImageService(repos)}
}
