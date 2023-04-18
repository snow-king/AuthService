package helpers

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Response struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *fiber.Ctx, response Response) error {
	if len(response.Message) > 0 {
		return c.Status(response.Status).JSON(map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		return c.Status(response.Status).JSON(map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
	return nil
}
