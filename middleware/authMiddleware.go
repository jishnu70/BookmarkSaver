package middleware

import (
	"bookmarksaver/initializers"
	"bookmarksaver/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(initializers.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		exp := int64(claims["exp"].(float64))

		if time.Now().Unix() > exp {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token expiread",
			})
			ctx.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		ctx.Set("userID", userID)

		var user models.User
		if err := initializers.DB.First(&user, userID).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			ctx.Abort()
			return
		}
		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
