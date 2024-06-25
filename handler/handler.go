package handler

import (
	"fiber-go/database"
	"fiber-go/model/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ReadAllUserndler(c *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	
	// err := database.DB.Find(&users).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }
	return c.JSON(users)
}

func Handler(c *fiber.Ctx) error {
	return c.SendString("Hello Bang")
}
