package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/models"
	"AuthService/app/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type response struct {
	Status string `json:"status"`
	Msg    string `json:"message,omitempty"`
}

func newResponse(status, msg string) *response {
	return &response{
		Status: status,
		Msg:    msg,
	}
}

func SignUp(c *gin.Context) {
	var inp models.LoginUser
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, newResponse(StatusError, err.Error()))
		return
	}
	//if err := server.AuthUseCase.SignUp(inp); err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(StatusError, err.Error()))
	//	return
	//}

	c.JSON(http.StatusOK, newResponse(StatusOk, "user created successfully"))
}

type signInResponse struct {
	*response
	Token string `json:"token,omitempty"`
}

func newSignInResponse(status, msg, token string) *signInResponse {
	return &signInResponse{
		&response{
			Status: status,
			Msg:    msg,
		},
		token,
	}
}

func SignIn(c *gin.Context) {
	var AuthUseCase = service.NewAuthorizer(
		viper.GetString("HASH_SALT"),
		[]byte(viper.GetString("SIGNING_KEY")),
		viper.GetDuration("TOKEN_TTL")*time.Second,
	)
	var inp models.LoginUser
	if err := c.Bind(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token, err := AuthUseCase.SignIn(inp)
	if err != nil {
		if err == errors.ErrInvalidAccessToken {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		if err == errors.ErrUserDoesNotHaveAccess {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		if err == errors.ErrUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(StatusError, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, newSignInResponse(StatusOk, "", token))
}
