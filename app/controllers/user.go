package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/service"
	"AuthService/app/structures"
	"github.com/gofiber/fiber/v2"
)

func AppendRole(c *fiber.Ctx) error {
	var inp structures.UserRole
	if err := c.BodyParser(&inp); err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	userS, err := service.NewUserService(inp.UserId)
	roles := userS.AppendRole(inp.RoleId)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			c.Status(fiber.StatusBadRequest)
			return nil
		}
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	return c.JSON(roles)
}
func AppendSocNetworks(c *fiber.Ctx) error {
	var inp structures.UserSocNetwork
	if err := c.BodyParser(&inp); err != nil {
		c.Status(fiber.StatusBadRequest)
		return nil
	}
	userS, err := service.NewUserService(inp.UserId)
	networks, err := userS.AddSocNetwork(inp.NetworkId, inp.NetworkNameId)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			c.Status(fiber.StatusBadRequest)
			return nil
		}
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(networks)
}
