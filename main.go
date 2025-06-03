package main

import (
	"bookmarksaver/controllers"
	"bookmarksaver/initializers"
	"bookmarksaver/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.GET("/root", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Backend is online",
		})
	})

	api := r.Group("/api")
	api.POST("/register", controllers.CreateUsers)
	api.POST("/login", controllers.Login)

	auth := r.Group("/auth/bookmark")
	auth.Use(middleware.RequireLogin())

	auth.GET("/", controllers.GetAllBookmark)
	auth.GET("/:id", controllers.GetBookMarkByID)
	auth.POST("/", controllers.AddBookMark)
	auth.PUT("/:id", controllers.UpdateBookMark)
	auth.DELETE("/:id", controllers.DeleteBookmark)

	r.Run()
}
