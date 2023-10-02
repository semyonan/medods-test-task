package app

import (
	"log"
	"medods-test-task/internal/config"
	"medods-test-task/internal/controllers"
	"medods-test-task/internal/database"

	_ "medods-test-task/internal/app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.apikey UsersAuth
// @in cookie
// @name access_token
func Run(configPath string) {

	cfg, err := config.Init(configPath)

	if (err != nil) {
		log.Fatal(err)
	}

	client, err := database.Init(cfg.Db.Url);

	if (err != nil) {
		log.Fatal(err)
	}

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.InitUserHandler(router, cfg, client)

	router.Run()
}