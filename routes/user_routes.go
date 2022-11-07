package routes

import (
	"github.com/concurrente-parcial/back/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	// app.Post("/login", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })
	// app.Post("/verification", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })

	app.Post("/user", controllers.CreateUser) //add this
}
