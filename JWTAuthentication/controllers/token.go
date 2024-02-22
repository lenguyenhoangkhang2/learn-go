package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenguyenhoangkhang2/go_authentication/auth"
	"github.com/lenguyenhoangkhang2/go_authentication/database"
	"github.com/lenguyenhoangkhang2/go_authentication/models"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @BasePath /api

// GenerateToken godoc
// @Router /token [post]
func GenerateToken(context *gin.Context) {
	var tokenReq TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&tokenReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Where("email = ?", user.Email).First(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(user.Password)

	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": credentialError.Error()})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
