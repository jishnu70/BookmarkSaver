package main

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.AuthInput{})
	initializers.DB.AutoMigrate(&models.Bookmark{})
	// initializers.DB.AutoMigrate(&models.CreateBookmarkInput{})
	initializers.DB.AutoMigrate(&models.Tag{})
	// initializers.DB.AutoMigrate(&models.UpdateBookMarkInput{})
}
