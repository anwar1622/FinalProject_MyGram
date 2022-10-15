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

func CreateSocialMedia(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		context.ShouldBindJSON(&SocialMedia)
	} else {
		context.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMedia(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}
	User := models.User{}

	UserID := uint(userData["id"].(float64))
	SocialMedia.UserID = UserID

	err := db.Where("user_id = ?", UserID).Find(&SocialMedia).Error
	errUser := db.Where("id = ?", UserID).Find(&User).Error
	if err != nil || errUser != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if SocialMedia.ID > 0 {
		context.JSON(http.StatusOK, gin.H{
			"social_media": gin.H{
				"id":               SocialMedia.ID,
				"name":             SocialMedia.Name,
				"social_media_url": SocialMedia.SocialMediaURL,
				"user_id":          SocialMedia.UserID,
				"created_at":       SocialMedia.CreatedAt,
				"updated_at":       SocialMedia.UpdatedAt,
				"User": gin.H{
					"id":       User.ID,
					"username": User.Username,
				},
			},
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "You don't have any data",
		})
	}
}

func UpdateSocialMedia(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)
	SocialMedia := models.SocialMedia{}
	socialMediaID, _ := strconv.Atoi(context.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		context.ShouldBindJSON(&SocialMedia)
	} else {
		context.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userId
	SocialMedia.ID = uint(socialMediaID)
	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(context.Param("socialMediaId"))
	userId := uint(userData["id"].(float64))

	SocialMedia.ID = uint(socialMediaId)
	SocialMedia.UserID = userId

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
