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

func CreateComment(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		context.ShouldBindJSON(&Comment)
	} else {
		context.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}
	User := models.User{}
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	Photo.UserID = userID
	User.ID = userID
	Comment.UserID = userID

	err := db.Where("user_id = ?", userID).Find(&Comment).Error
	errUser := db.Where("id = ?", userID).Find(&User).Error
	errPhoto := db.Where("user_id= ?", userID).Find(&Photo).Error

	if err != nil || errUser != nil || errPhoto != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
		"created_at": Comment.CreatedAt,
		"User": gin.H{
			"id":       User.ID,
			"email":    User.Email,
			"username": User.Username,
		},
		"Photo": gin.H{
			"id":        Photo.ID,
			"title":     Photo.Title,
			"caption":   Photo.Caption,
			"photo_url": Photo.PhotoURL,
			"user_id":   Photo.UserID,
		},
	})
}

func UpdateComment(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(context.Param("commentId"))

	userId := uint(userData["id"].(float64))
	if contentType == appJSON {
		context.ShouldBindJSON(&Comment)
	} else {
		context.ShouldBind((&Comment))
	}

	Comment.ID = uint(commentId)
	Comment.UserID = uint(userId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, Comment)
}

func DeleteComment(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(context.Param("commentId"))
	userId := uint(userData["id"].(float64))

	Comment.ID = uint(commentId)
	Comment.UserID = userId

	err := db.Model(&Comment).Where("id = ?", commentId).Delete(models.Comment{}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
