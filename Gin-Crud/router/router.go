package router

import (
	"TrainingGo/Gin-Crud/handler"
	"TrainingGo/Gin-Crud/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/", handler.RootHandler) // localhost:/

	// // Tambahkan middleware AuthMiddleware ke rute yang memerlukan autentikasi
	privateEndPoint := r.Group("/private") // localhost:/private
	privateEndPoint.Use(middleware.AuthMiddleware())
	{
		privateEndPoint.POST("/", handler.PostHandler) // localhost:/private
	}

	r.POST("/users/add", handler.CreateUsers)
	r.GET("/users", handler.GetAllUsers)
	// r.GET("/users/:id", handler.GetUserByID)
	r.GET("/users/:name", handler.GetUserByName)

	/*
		usersPublicEndpoint := r.Group("/users")
		usersPublicEndpoint.GET("/:id", handler.GetUser)
		usersPublicEndpoint.GET("/", handler.GetAllUsers)

		usersPrivateEndpoint := r.Group("/users")
		usersPrivateEndpoint.Use(middleware.AuthMiddleware())
		usersPrivateEndpoint.POST("/", handler.CreateUser)
		usersPrivateEndpoint.PUT("/:id", handler.UpdateUser)
		usersPrivateEndpoint.DELETE("/:id", handler.DeleteUser)
	*/
}
