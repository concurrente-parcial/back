package main

import (
	"github.com/concurrente-parcial/back/configs"
	"github.com/concurrente-parcial/back/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":6000")
}
