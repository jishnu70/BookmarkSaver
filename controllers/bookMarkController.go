package controllers

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddBookMark(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	var inputData models.CreateBookmarkInput

	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
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
		Title:  inputData.Title,
		URL:    inputData.URL,
		Tags:   tags,
		UserID: userID,
	}

	if err := initializers.DB.Create(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bookmark"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"bookMark": bookMark,
	})
}

func GetAllBookmark(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	var bookMarks []models.Bookmark

	if err := initializers.DB.Preload("Tags").Where("user_id=?", userID).Find(&bookMarks).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bookmarks": bookMarks,
	})
}

func GetBookMarkByID(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	id := ctx.Param("id")

	var bookMark models.Bookmark

	if err := initializers.DB.Preload("Tags").Where("id=? AND user_id=?", id, userID).First(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bookMark": bookMark,
	})
}

func UpdateBookMark(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	id := ctx.Param("id")

	var bookMark models.Bookmark

	if err := initializers.DB.Preload("Tags").Where("id=? AND user_id=?", id, userID).First(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var inputData models.UpdateBookMarkInput
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if inputData.Title != nil {
		bookMark.Title = *inputData.Title
	}

	if inputData.URL != nil {
		bookMark.URL = *inputData.URL
	}

	if inputData.Tags != nil {
		var tags []models.Tag
		for _, tagName := range inputData.Tags {
			var tag models.Tag
			if err := initializers.DB.Where("name=?", tagName).First(&tag).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					tag = models.Tag{Name: tagName}
					initializers.DB.Create(&tag)
				} else {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Tag lookup failed"})
					return
				}
			}
			tags = append(tags, tag)
		}
		initializers.DB.Model(&bookMark).Association("Tags").Replace(tags)
	}

	if err := initializers.DB.Save(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bookmark"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"bookMark": bookMark})
}

func DeleteBookmark(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	id := ctx.Param("id")

	var bookMark models.Bookmark

	if err := initializers.DB.Preload("Tags").Where("id=? AND user_id=?", id, userID).First(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := initializers.DB.Delete(&bookMark).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bookmark deleted",
	})
}
