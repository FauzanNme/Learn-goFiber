package routes

import (
	"fiber-go/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// PATH
	r.Get("/", handler.Handler)
	r.Get("/user", handler.ReadAllUserndler)

	// Bentuk JSON
	// //nama ctx bebas
	// r.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.JSON(fiber.Map{
	// 		"hello": "world",
	// 	})
	// })
}
