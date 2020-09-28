package controller

import (
	"net/http"

	"video-api/services"

	"github.com/gin-gonic/gin"
)

func CloudflareSearchVideo(c *gin.Context) {

	videoName := c.Param("videoName")

	var cloudflareService services.CloudflareService
	if videoSearchResponse := cloudflareService.SearchVideo(videoName); cloudflareService.Error != nil {
		c.JSON(http.StatusBadRequest, cloudflareService.Error)
		return
	} else {
		c.JSON(http.StatusOK, videoSearchResponse)
		return
	}
}

func CloudflareUploadVideo(c *gin.Context) {

	_, header, err := c.Request.FormFile("videoName")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var cloudflareService services.CloudflareService
	if uploadReturnModel := cloudflareService.UploadVideo(header); cloudflareService.Error != nil {
		c.JSON(http.StatusBadRequest, cloudflareService.Error)
		return
	} else {
		c.JSON(http.StatusOK, uploadReturnModel)
		return
	}
}

func CloudflareSignedURL(c *gin.Context) {

	videoUID := c.Param("videoUID")

	var cloudflareService services.CloudflareService
	if url := cloudflareService.GetSignedURL(videoUID); cloudflareService.Error != nil {
		c.JSON(http.StatusBadRequest, cloudflareService.Error)
		return
	} else {
		c.JSON(http.StatusOK, url)
		return
	}
}
