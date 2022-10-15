package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePhoto(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	if contentType == appJSON {
		context.ShouldBindJSON(&Photo)
	} else {
		context.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetPhoto(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	User := models.User{}

	userID := uint(userData["id"].(float64))
	Photo.UserID = userID
	err := db.Where("user_id = ?", userID).Find(&Photo).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error

	if errUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if Photo.ID > 0 {
		context.JSON(http.StatusOK, gin.H{
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoURL,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
			"updated_at": Photo.UpdatedAt,
			"User": gin.H{
				"email":    User.Email,
				"username": User.Username,
			},
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "You don't have any photo",
		})
	}
}

func UpdatePhoto(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(context.Param("photoId"))
	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		context.ShouldBindJSON(&Photo)
	} else {
		context.ShouldBind(&Photo)
	}
	Photo.UserID = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(context.Param("photoId"))
	userId := uint(userData["id"].(float64))
	Photo.ID = uint(photoId)
	Photo.UserID = userId

	err := db.Model(&Photo).Where("id = ?", photoId).Delete(models.Photo{}).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
