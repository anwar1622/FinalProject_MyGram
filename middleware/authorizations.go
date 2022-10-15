package middleware

import (
	"MyGram/database"
	"MyGram/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PhotoAutorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(context.Param("photoId"))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}
		err = db.Select("user_id").First(&Photo, uint(photoId)).Error
		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}
		if Photo.UserID != userID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		context.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(context.Param("commentId"))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid parameter",
			})
			return
		}
		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}
		err = db.Select("user_id").First(&Comment, uint(commentId)).Error
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}
		if Comment.UserID != userID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		context.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		db := database.GetDB()
		socialMediaId, err := strconv.Atoi(context.Param("socialMediaId"))
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		SocialMedia := models.SocialMedia{}
		err = db.Select("user_id").First(&SocialMedia, uint(socialMediaId)).Error

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Data not found",
				"message": "Data doesn't exist",
			})
			return
		}
		if SocialMedia.UserID != userID {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		context.Next()
	}
}
