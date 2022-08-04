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

	router.LoadHTMLFiles("D:\\image-softcery\\templates\\upload.html")

	images := router.Group("/images")
	{
		images.GET("/upload", h.uploadTemplate)
		images.POST("/upload", h.imageUpload)

		images.GET("/download", h.downloadTemplate)
		images.POST("/download", h.imageDownload)
	}
	return router
}