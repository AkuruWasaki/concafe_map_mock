package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/con-cafe", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "concafe map",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
