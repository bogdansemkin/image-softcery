package handlers

import (
	"github.com/bogdansemkin/image-softcery/pkg/rabbit"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// @Summary     imageUpload
// @Tags        images
// @Description upload image on server
// @ID          upload-image
// @Accept      file
// @Produce     json
// @Param       input   body     file true "file"
// @Success     200     {string} string
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /images/upload [post]

func (h *Handler) imageUpload(c *gin.Context) {
	mq := rabbit.MQ{}
	file, err := c.FormFile("imageFile")
	if err != nil {
		logrus.Errorf("error on controller upload, %s", err)
	}

	tempFile, err := ioutil.TempFile("D:\\image-softcery\\templates\\img", "upload-*.png")
	if err != nil {
		logrus.Errorf("Error during creating temp image, %s", err)
	}

	fileNew, _ := file.Open()

	fileBytes, err := ioutil.ReadAll(fileNew)
	if err != nil {
		logrus.Errorf("Error during readAll temp image, %s", err)
	}

	tempFile.Write(fileBytes)

	mq.Producer(tempFile.Name())
	image, seventyFiveImage, halfImage, partImage := mq.Consumer()

	//cache.Set(fmt.Sprintf("%s:%s", "upload", tempFile.Name()), "", cache.OneDayTtl)

	id, err := h.service.Upload(image, seventyFiveImage, halfImage, partImage)
	if err != nil{
		logrus.Errorf("Error during image uploading, %s", err)
	}

	c.String(http.StatusOK, "Successfully Uploaded File with id %d \n", id)
}

// @Summary     imageDownload
// @Tags        images
// @Description download image from server
// @ID          download-image
// @Accept      int
// @Produce     file
// @Param       input   body     int true "id"
// @Success     200     {file}   file
// @Failure     400,404 {object} errorResponse
// @Failure     500     {object} errorResponse
// @Failure     default {object} errorResponse
// @Router      /images/download/{id} [get]

func (h *Handler) imageDownload(c *gin.Context) {
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
	} else {
		c.File(image.Path)
	}
}
