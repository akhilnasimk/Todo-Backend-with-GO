package middlewares

import "github.com/gofiber/fiber/v2"

type U_auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Authmiddle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Next()
		return nil
	}
}
