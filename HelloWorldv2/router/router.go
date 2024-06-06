package router

import (
	"HelloWorldv2/HelloWorldv2/middleware"
	"HelloWorldv2/Helloworldv2/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/", handler.RootHandler) // localhost:/

	// // Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndPoint := r.Group("/private") // localhost/private
	privateEndPoint.Use(middleware.AuthMiddleware())
	{
		privateEndPoint.POST("/", handler.PostHandler) // localhost/private
	}
}
