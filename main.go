package main

import (
	"fiber-go/database"
	"fiber-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INIT database
	database.DatabaseInit()
	app := fiber.New()

	// INIT ROUTE
	routes.RouteInit(app)

	app.Listen(":3000")
}
