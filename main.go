package main

import (
	"bookmarksaver/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConntectDB()
}

func main() {
	r := gin.Default()

	r.GET("/root", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Backend is online",
		})
	})

	r.Run()
}
