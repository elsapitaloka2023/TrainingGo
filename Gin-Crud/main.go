package main

import (
	"TrainingGo/Gin-Crud/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//definisi route
	router.SetupRouter(r)

	//run di server 8080
	r.Run(": 8080")

}
