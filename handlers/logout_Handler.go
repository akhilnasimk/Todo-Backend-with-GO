package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogOut(c *fiber.Ctx) error {
	token := c.Cookies("access_token")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{"message": "User not logged in"})
	}
	accessCookie := fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&accessCookie)

	//setting the reference cookie
	RefreshCookie := fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&RefreshCookie) //setted here

	return c.Status(200).JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}
