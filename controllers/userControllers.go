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

var appJSON = "application/json"

func UserRegister(context *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	if appJSON == helpers.GetContentType(context) {
		context.ShouldBindJSON(&User)
	} else {
		context.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func UserLogin(context *gin.Context) {
	db := database.GetDB()
	contentType := context.Request.Header.Get("Content-Type")
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		context.ShouldBindJSON(&User)
	} else {
		context.ShouldBind(&User)
	}
	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email/Password",
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)
	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(context)
	User := models.User{}

	paramId, _ := strconv.Atoi(context.Param("userId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		context.ShouldBindJSON(&User)
	} else {
		context.ShouldBind(&User)
	}

	User.ID = userId

	err := db.Model(&User).Where("id = ?", paramId).Updates(models.User{Username: User.Username, Email: User.Email, Password: User.Password, Age: User.Age}).Error
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func DeleteUser(context *gin.Context) {
	db := database.GetDB()
	userData := context.MustGet("userData").(jwt.MapClaims)
	User := models.User{}
	userID := uint(userData["id"].(float64))
	User.ID = userID
	err := db.Model(&User).Where("id = ?", userID).Delete(models.User{}).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
