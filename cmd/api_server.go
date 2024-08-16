package cmd

import "github.com/gin-gonic/gin"

func StartApiServer() {
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run(":8001")
}
