package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenguyenhoangkhang2/go_authentication/database"
	"github.com/lenguyenhoangkhang2/go_authentication/models"
)

// @BasePath /api

// RegisterUser godoc
// @Param request body models.User true "query params"
// @Router /user/register [post]
func RegisterUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HassPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
