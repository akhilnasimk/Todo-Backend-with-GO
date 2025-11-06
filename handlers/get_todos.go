package handlers

import (
	"todo/services"

	"github.com/gofiber/fiber/v2"
)

func GetAlltodo(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(string)
	todos := services.GetTodos(user_id)
	return c.Status(200).JSON(fiber.Map{
		"todos": todos,
	})
}
