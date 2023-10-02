package controllers

import (
	"medods-test-task/internal/utils/tokens"
	"medods-test-task/internal/config"
	"medods-test-task/internal/utils/hash"
	"medods-test-task/internal/services"
	"medods-test-task/internal/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitUserHandler(router *gin.Engine, cfg *config.Config, client *mongo.Client) {
	userRepo := repositories.NewUsersRepo(client.Database(cfg.Db.Name))
	sessionRepo := repositories.NewSessionRepo(client.Database(cfg.Db.Name))
	hasher := hash.NewBcryptHasher()
	tokenManager, _ := tokens.NewManager(cfg.Auth.Secret)

	service := services.NewUsersService(*userRepo, *sessionRepo, hasher, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	
	handler := NewUserHandler(service)

	// task routes
	router.GET("/refresh", handler.SignOut)
	router.GET("/id-signin/:id", handler.IdSignIn)

	// help routes
	router.GET("/signin/", handler.SignIn)
	router.POST("/signup", handler.SignUp)
	router.POST("/signout", handler.SignOut)
}

func InitUserTestHandler(router *gin.Engine, service services.UsersServiceInterface) {
	
	handler := NewUserHandler(service)

	// task routes
	router.GET("/refresh", handler.SignOut)
	router.GET("/id-signin/:id", handler.IdSignIn)

	// help routes
	router.GET("/signin", handler.SignIn)
	router.POST("/signup", handler.SignUp)
	router.POST("/signout", handler.SignOut)
}