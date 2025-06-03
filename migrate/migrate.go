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
}
