package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/models"
	"AuthService/app/service"
	"github.com/gofiber/fiber/v2"
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

func SignUp(c *fiber.Ctx) error {
	var inp models.LoginUser
	if err := c.BodyParser(&inp); err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	//if err := server.AuthUseCase.SignUp(inp); err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, newResponse(StatusError, err.Error()))
	//	return
	//}
	return c.Status(fiber.StatusCreated).JSON(newResponse(StatusOk, "user created successfully"))
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

func SignIn(c *fiber.Ctx) error {
	var AuthUseCase = service.NewAuthorizer(
		viper.GetString("HASH_SALT"),
		[]byte(viper.GetString("SIGNING_KEY")),
		viper.GetDuration("TOKEN_TTL")*time.Second,
	)
	var inp models.LoginUser
	if err := c.BodyParser(&inp); err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	token, err := AuthUseCase.SignIn(inp)
	if err != nil {
		if err == errors.ErrInvalidAccessToken {
			return c.Status(fiber.StatusForbidden).JSON(newSignInResponse(StatusError, err.Error(), ""))
		}
		if err == errors.ErrUserDoesNotHaveAccess {
			return c.Status(fiber.StatusForbidden).JSON(newSignInResponse(StatusError, err.Error(), ""))
		}
		if err == errors.ErrUserDoesNotExist {
			return c.Status(fiber.StatusBadRequest).JSON(newSignInResponse(StatusError, err.Error(), ""))
		}
		return c.Status(http.StatusInternalServerError).JSON(newSignInResponse(StatusError, err.Error(), ""))
	}
	return c.Status(http.StatusOK).JSON(newSignInResponse(StatusOk, "", token))
}
