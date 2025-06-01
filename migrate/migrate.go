package main

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConntectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
