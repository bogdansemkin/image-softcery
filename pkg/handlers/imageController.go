package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func (h *Handler) imageUpload(ctx *gin.Context){
	file, err := ctx.FormFile("imageFile")
	if err != nil {
		logrus.Errorf("error on controller upload, %s", err)
	}

	fmt.Printf("Uploaded file: %s", file.Filename)
	fmt.Printf("Size of uploaded file: %d", file.Size)
	fmt.Printf("Header of uploaded file: %s", file.Header)

	tempFile, err := ioutil.TempFile("templates\\img", "upload-*.png")
	if err != nil{
		logrus.Errorf("Error during creating temp image, %s", err)
	}

	fileNew, _ := file.Open()

	fileBytes, err := ioutil.ReadAll(fileNew)
	if err != nil{
		logrus.Errorf("Error during readAll temp image, %s", err)
	}

	tempFile.Write(fileBytes)

	h.service.Image.Upload(tempFile.Name())
	ctx.String(http.StatusOK, "Successfully Uploaded File\n")
}

func (h *Handler) imageDownload(ctx *gin.Context){

}

func (h *Handler) uploadTemplate(ctx *gin.Context){
	ctx.HTML(http.StatusOK, "upload.html", gin.H{
		"title":    "Upload image",
	})
}

func (h *Handler) downloadTemplate(ctx *gin.Context){
	id := ctx.Param("id")
	image, err := h.service.Download(id)
	if err != nil{
		logrus.Errorf("Error on download template, %s", err)
	}
	testString := strings.TrimPrefix(image.Path, "templates\\")
	ctx.HTML(http.StatusOK, "download.html", gin.H{
		"Path": template.URL(testString),
	})
}



