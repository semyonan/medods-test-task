package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"medods-test-task/internal/DAO"
	"medods-test-task/internal/controllers/mock_services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUserHandler_SignUp(t *testing.T) {
	var bodyMsg map[string]string
	mockCtrl := gomock.NewController(t)

	Convey("Given POST /signup", t, func() {
		w := httptest.NewRecorder()
		data := DAO.SignUpRequest{
			Email: "email",
			Password: "password",
		}
		reqData,_ := json.Marshal(&data)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewReader(reqData))
    

		Convey("When there is no user with the same email", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			mock.EXPECT().SignUp(gomock.Eq(context.Background()), gomock.Eq(data)).Return(nil)

			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)

			Convey("Then 200 OK is returned", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})
		})

		Convey("When there is user with the same email", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			mock.EXPECT().SignUp(gomock.Eq(context.Background()), gomock.Eq(data)).Return(errors.New("you already have account"))

			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)

			Convey("Then 404 Bad Request is returned", func() {
				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})
		})
	})
}

func TestAuthHandler_SignIn(t *testing.T) {
	var bodyMsg DAO.Tokens
	mockCtrl := gomock.NewController(t)

	Convey("Given GET /signin", t, func() {
		w := httptest.NewRecorder()
		data := DAO.SignInRequest{
			Email: "email",
			Password: "password",
		}
		reqData,_ := json.Marshal(&data)
		req, _ := http.NewRequest(http.MethodGet, "/signin", bytes.NewReader(reqData))

		Convey("When there is no", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			tokens := DAO.Tokens{}
			mock.EXPECT().SignIn(gomock.Eq(context.Background()), gomock.Eq(data)).Return(tokens, errors.New("there is no user with email"))
	
			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)
	
			Convey("Then 400 Unauthorized is returned", func() {
				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})
	
			Convey("And response body has the expected error message", func() {
				So(bodyMsg, ShouldEqual, tokens)
			})
		})	

		Convey("When there is user", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			tokens := DAO.Tokens{
				AccessToken: "access_token",
				AccessTokenTTL: 10,
				RefreshToken: "refresh_token",
				RefreshTokenTTL: 30,
			}
			mock.EXPECT().SignIn(gomock.Eq(context.Background()), gomock.Eq(data)).Return(tokens, nil)

			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)

			Convey("Then 200 OK is returned", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})

			Convey("And it is the expected tokens", func() {
				So(bodyMsg, ShouldEqual, tokens)
			})
		})
	})
}

func TestAuthHandler_IdSignIn(t *testing.T) {
	var bodyMsg DAO.Tokens
	mockCtrl := gomock.NewController(t)

	Convey("Given GET /id-signin/{id}", t, func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/id-signin/someid", nil)

		Convey("When there is no", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			tokens := DAO.Tokens{}
			mock.EXPECT().IdSignIn(gomock.Eq(context.Background()), gomock.Eq("someid")).Return(tokens, errors.New("there is no user with email"))
	
			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)
	
			Convey("Then 400 Unauthorized is returned", func() {
				So(w.Code, ShouldEqual, http.StatusBadRequest)
			})
	
			Convey("And response body has the expected error message", func() {
				So(bodyMsg, ShouldEqual, tokens)
			})
		})	

		Convey("When there is user", func() {
			gin.SetMode(gin.TestMode)
			router := gin.New()
			mock :=  mock_services.NewMockUsersServiceInterface(mockCtrl)
			InitUserTestHandler(router, mock)
			tokens := DAO.Tokens{
				AccessToken: "access_token",
				AccessTokenTTL: 10,
				RefreshToken: "refresh_token",
				RefreshTokenTTL: 30,
			}
			mock.EXPECT().IdSignIn(gomock.Eq(context.Background()), gomock.Eq("someid")).Return(tokens, nil)

			router.ServeHTTP(w, req)
			_ = json.Unmarshal(w.Body.Bytes(), &bodyMsg)

			Convey("Then 200 OK is returned", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})

			Convey("And it is the expected tokens", func() {
				So(bodyMsg, ShouldEqual, tokens)
			})
		})
	})
}
