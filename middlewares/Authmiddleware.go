package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authmiddle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("access_token")
		fmt.Println(tokenString)
		if tokenString == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "user not logged int no jwt ",
			})
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			jwtSecret := os.Getenv("JWT_SECRET")
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}
		claim := token.Claims.(jwt.MapClaims)
		User_id := claim["user_id"].(string)
		// fmt.Println("user_id is ", User_id)
		c.Locals("user_id", User_id)
		return c.Next()
	}
}
