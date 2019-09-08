package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snail/app"
	"snail/pkg/e"
	"snail/pkg/logging"
	"snail/pkg/upload"
)

func UploadImage(c *gin.Context) {
	at := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		at.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		at.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		at.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		at.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		at.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	at.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"url":      upload.GetImageFullUrl(imageName),
		"save_url": savePath + imageName,
	})
}
