package controllers

import (
	"medods-test-task/internal/DAO"
	"medods-test-task/internal/services"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UsersServiceInterface
}

func NewUserHandler(service services.UsersServiceInterface) *UserHandler{
	return &UserHandler{service: service}
}


// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body DAO.SignUpRequest true "sign up info"
// @Router /signup [post]
func (h *UserHandler) SignUp(c *gin.Context)  {
	var req DAO.SignUpRequest

	err := c.BindJSON(&req)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	err = h.service.SignUp(c.Request.Context(), req)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	c.Status(http.StatusOK)	
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body DAO.SignInRequest true "sign up info"
// @Success 200 {object} DAO.Tokens
// @Router /signin [get]
func (h *UserHandler) SignIn(c *gin.Context) {
	var req DAO.SignInRequest
	err := c.BindJSON(&req)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	tokens, err := h.service.SignIn(c.Request.Context(), req)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	c.JSON(http.StatusOK, tokens)	
	c.SetCookie("access_token", tokens.AccessToken,  int(tokens.AccessTokenTTL.Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(tokens.RefreshTokenTTL.Seconds()), "/", "localhost", false, true)
}

// @Summary User SignOut
// @Tags users-auth
// @Description user sign out
// @ModuleID userSignOut
// @Accept  json
// @Produce  json
// @Success 200 {object} DAO.Response
// @Router /signout [post]
func (h *UserHandler) SignOut(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.Status(http.StatusOK)
}

// @Summary User IdSignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userIdSignIn
// @Accept  json
// @Produce  json
// @Param id path string true "user id"
// @Success 200 {object} DAO.Tokens
// @Router /id-signin/{id} [get]
func (h *UserHandler) IdSignIn(c *gin.Context) {
	id := c.Param("id")

	tokens, err := h.service.IdSignIn(c.Request.Context(), id)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	c.JSON(http.StatusOK, tokens)	
	c.SetCookie("access_token", tokens.AccessToken,  int(tokens.AccessTokenTTL.Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(tokens.RefreshTokenTTL.Seconds()), "/", "localhost", false, true)
}

// @Summary User Refresh
// @Tags users-auth
// @Description user refresh
// @ModuleID userRefresh
// @Accept  json
// @Produce  json
// @Success 200 {object} DAO.Tokens
// @Router /refresh [get]
func (h *UserHandler) Refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	tokens, err := h.service.RefreshTokens(c.Request.Context(), cookie)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())	
		return
	}

	c.JSON(http.StatusOK, tokens)	
	c.SetCookie("access_token", tokens.AccessToken,  int(tokens.AccessTokenTTL.Seconds()), "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokens.RefreshToken, int(tokens.RefreshTokenTTL.Seconds()), "/", "localhost", false, true)
}