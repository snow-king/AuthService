package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/service/SocialService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func choseSocial(name string) SocialService.SocService {
	var network SocialService.SocService
	switch name {
	case "vk":
		network = SocialService.NewVkService()
	default:
		return nil
	}
	return network
}
func Index(c *fiber.Ctx) error {
	socNetworkName := c.Params("service")
	network := choseSocial(socNetworkName)
	if network != nil {
		url := network.Index()
		return c.Redirect(url, fiber.StatusMovedPermanently)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(newSignInResponse(StatusError, "не верный сервис авторизации", ""))
	}
}
func Callback(c *fiber.Ctx) error {
	socNetworkName := c.Params("service")
	network := choseSocial(socNetworkName)
	authCode := c.Query("code")
	token, err := network.Callback(authCode)
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
