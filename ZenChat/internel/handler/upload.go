package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"net/http"
)

type UploadHandler struct{}

var NewUploadHandler = &UploadHandler{}

// Upload 图片、语音文件上传
func (u *UploadHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("filename")
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	url, err := itools.QiNiuFileUpload(c.Request(), file.Filename)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "upload success", url))
}
