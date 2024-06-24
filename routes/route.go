package routes

import (
	"fiber-go/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// PATH
	// nama c karo r bebas
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	r.Get("/user", handler.ReadUserHandler)

	// Bentuk JSON
	// r.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.JSON(fiber.Map{
	// 		"hello": "world",
	// 	})
	// })
}
