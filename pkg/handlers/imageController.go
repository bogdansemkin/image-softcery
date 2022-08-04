package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) imageUpload(ctx *gin.Context){

}

func (h *Handler) imageDownload(ctx *gin.Context){

}

func (h *Handler) uploadTemplate(ctx *gin.Context){
	ctx.HTML(http.StatusOK, "upload.html", gin.H{
		"title":    "Upload image",
	})
}

func (h *Handler) downloadTemplate(ctx *gin.Context){

}



