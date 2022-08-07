package handlers

import (
	"github.com/bogdansemkin/image-softcery/pkg/services"
	"github.com/gin-gonic/gin"

	 "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	docs "github.com/bogdansemkin/image-softcery/docs"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath="/pkg/hanlders/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	images := router.Group("/images")
	{
		images.POST("/upload", h.imageUpload)
		images.GET("/download/:id", h.imageDownload)
	}
	return router
}
