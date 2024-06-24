package handler

import "github.com/gofiber/fiber/v2"

func ReadUserHandler(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"data":"user",
		})
}    