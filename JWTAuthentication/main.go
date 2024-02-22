package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/lenguyenhoangkhang2/go_authentication/controllers"
	"github.com/lenguyenhoangkhang2/go_authentication/database"
	"github.com/lenguyenhoangkhang2/go_authentication/docs"
	"github.com/lenguyenhoangkhang2/go_authentication/middlewares"
)

var (
	DB_USER     = "postgres"
	DB_NAME     = "golang_jwt"
	DB_PASSWORD = "khang"
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	database.Connect(dsn)
	database.Migrate()

	router := initRouter()

	router.Run(":8080")
}

func initRouter() *gin.Engine {
	gin.ForceConsoleColor()

	router := gin.Default()

	basePath := "/api"

	docs.SwaggerInfo.BasePath = basePath

	api := router.Group(basePath)

	{
		api.POST("/token", controllers.GenerateToken)

		api.POST("/user/register", controllers.RegisterUser)

		secured := api.Group("/secured").Use(middlewares.Auth())

		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
