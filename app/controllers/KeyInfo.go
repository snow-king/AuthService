package controllers

import (
	"AuthService/app/service"
	"github.com/gofiber/fiber/v2"
)

func JWK(c *fiber.Ctx) error {
	JwtInfo := new(service.JWTInfo)
	jwk := JwtInfo.InitJWK()
	return c.JSON(jwk)
}
