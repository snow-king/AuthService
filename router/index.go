package router

import (
	"AuthService/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router) {
	router.Post("/sign-up", controllers.SignUp)
	router.Post("/sign-in", controllers.SignIn)
	router.Get("/O2Auth/:service/index", controllers.Index)
	router.Get("/O2Auth/:service/callback", controllers.Callback)
	router.Post("/user/appendRole", controllers.AppendRole)
	router.Post("/user/appendNetwork", controllers.AppendSocNetworks)
	router.Get("/jwk", controllers.JWK)
	router.Get("/health", func(context *fiber.Ctx) error {
		return context.JSON(fiber.Map{"response": "It's Alive! Alive!!!!"})
	})
}
