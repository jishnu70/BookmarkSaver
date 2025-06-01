package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/root", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Backend is online",
		})
	})

	r.Run()
}
