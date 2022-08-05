package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func (h *Handler) imageUpload(c *gin.Context){
	file, err := c.FormFile("imageFile")
	if err != nil {
		logrus.Errorf("error on controller upload, %s", err)
	}

	tempFile, err := ioutil.TempFile("D:\\image-softcery\\templates\\img", "upload-*.png")
	if err != nil{
		logrus.Errorf("Error during creating temp image, %s", err)
	}

	fileNew, _ := file.Open()

	fileBytes, err := ioutil.ReadAll(fileNew)
	if err != nil{
		logrus.Errorf("Error during readAll temp image, %s", err)
	}

	tempFile.Write(fileBytes)

	_ , err = h.service.Image.Upload(tempFile.Name(), imageResize(tempFile.Name()), imageHalfResize(tempFile.Name()), imageFullResize(tempFile.Name()))
	if err != nil {
		logrus.Errorf("Error on uploading image, %s", err)
	}

	c.String(http.StatusOK, "Successfully Uploaded File\n")
}

func (h *Handler) imageDownload(c *gin.Context){
	id := c.Param("id")
	quality := c.Query("quality")

	image, err := h.service.Image.Download(id)
	if err != nil {
		logrus.Errorf("Error during downloading image... %s", err)
	}
	if quality != "" {
		switch quality {
		case "100":
			c.File(image.Path)
		case "75":
			c.File(image.SeventyFivePath)
		case "50":
			c.File(image.HalfPath)
		case "25":
			c.File(image.TwentyFivePath)
		}
	} else{
		c.File(image.Path)
	}
}