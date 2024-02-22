package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func uploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": http.ErrMissingFile.Error()})
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))

	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt

	filePath := "http://localhost:8000/images/public/" + filename

	out, err := os.Create("public/images/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"filepath": filePath})
}

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "How to Upload Single and Multiple Files in Golang"})
	})

	r.POST("/upload-image", uploadImage)

	err := r.Run(":8080")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
