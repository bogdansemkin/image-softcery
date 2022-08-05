package handlers

import (
	"github.com/gin-gonic/gin"
	"image-softcery/pkg/services"
)

type Handler struct{
	service *services.Service
}

func NewHandler(service *services.Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.Default()

	images := router.Group("/images")
	{
		images.POST("/upload", h.imageUpload)
		images.GET("/download/:id", h.imageDownload)
	}
	return router
}