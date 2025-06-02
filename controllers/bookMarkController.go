package controllers

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddBookMark(ctx *gin.Context) {
	var inputData models.CreateBookmarkInput

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	// prepare for tags
	var tags []models.Tag
	for _, tagName := range inputData.Tags {
		var tag models.Tag
		if err := initializers.DB.Where("name=?", tagName).First(&tag).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tag = models.Tag{Name: tagName}
				initializers.DB.Create(&tag)
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "DB error when finding tag"})
				return
			}
		}
		tags = append(tags, tag)
	}

	bookMark := models.Bookmark{
		Title: inputData.Title,
		URL:   inputData.URL,
		Tags:  tags,
	}

	if err := initializers.DB.Create(&bookMark); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bookmark"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"bookMark": bookMark,
	})
}

func GetAllBookmark(ctx *gin.Context) {
	var bookMarks []models.Bookmark

	if err := initializers.DB.Preload("Tags").Find(&bookMarks).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bookmarks": bookMarks,
	})
}
