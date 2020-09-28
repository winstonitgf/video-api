package routers

import (
	"video-api/controller"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/cloudflare/video/search/:videoName", controller.CloudflareSearchVideo)
			v1.POST("/cloudflare/video/upload", controller.CloudflareUploadVideo)
			v1.GET("/cloudflare/video/sign/:videoUID", controller.CloudflareSignedURL)
		}
	}
	r.Run()
}
