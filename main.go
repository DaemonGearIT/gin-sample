package main

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hallo world",
		})
	})

	color.Green("Server started at --> http://127.0.0.1:9000/")
	r.Run(":9000")
}
