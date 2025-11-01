package handlers

import "github.com/gofiber/fiber/v2"

func GetAlltodo(c *fiber.Ctx) error {
	tocken := c.Cookies("access_token")

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "here all the todos",
		"jwt":     tocken,
	})
}
