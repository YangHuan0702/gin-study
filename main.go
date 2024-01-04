package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "haloSky"})
	})
	r.Run(":8080")
}
