package controllers

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUsers(ctx *gin.Context) {
	var authInput models.AuthInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"User": user,
	})
}

func Login(ctx *gin.Context) {
	var authInput models.AuthInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetUserProfile(ctx *gin.Context) {
	user, _ := ctx.Get("currentUser")

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
