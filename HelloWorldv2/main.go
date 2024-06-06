package main

import (
	"HelloWorldv2/Helloworldv2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//definisi route
	router.SetupRouter(r)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello",
	// 	})
	// })

	// r.POST("/post", func(c *gin.Context) {
	// 	var json struct {
	// 		Message string `json:"message"`
	// 	}

	// 	if err := c.ShouldBindJSON(&json); err == nil {
	// 		c.JSON(200, gin.H{"message": json.Message})
	// 	} else {
	// 		c.JSON(400, gin.H{"error": err.Error()})
	// 	}
	// })

	//run di server 8080
	r.Run(": 8080")

}
