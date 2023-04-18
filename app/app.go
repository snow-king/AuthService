package server

import (
	"AuthService/app/models"
	router2 "AuthService/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string) {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "AuthorizationService",
		AppName:       "AuthorizationService IrGUPS v 1.0",
	})
	models.ConnectDatabase()
	// Endpoints
	app.Use(logger.New(), recover.New())
	app.Route(
		"/auth",
		router2.RegisterHTTPEndpoints,
		"auth.",
	)
	//router2.RegisterHTTPEndpoints(api)
	//api.Post("/sign-up", controllers.SignUp)
	//api.Post("/sign-in", controllers.SignIn)
	//api.Get("/O2Auth/:service/index", controllers.Index)
	//api.Get("/O2Auth/:service/callback", controllers.Callback)
	//api.Post("/user/appendRole", controllers.AppendRole)
	//api.Post("/user/appendNetwork", controllers.AppendSocNetworks)
	//api.Get("/jwk", controllers.JWK)
	//api.Get("/health", func(context *fiber.Ctx) error {
	//	return context.JSON(fiber.Map{"response": "It's Alive! Alive!!!!"})
	//})
	// HTTP Server
	log.Fatal(app.Listen(":" + port))
}
