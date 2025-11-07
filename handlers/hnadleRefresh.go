package handlers

import (
	"time"
	"todo/db"
	"todo/services"

	"github.com/gofiber/fiber/v2"
)

func RefreshHandler(c *fiber.Ctx) error {
	refreshtocken := c.Cookies("refresh_token")

	if refreshtocken == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Refresh token missing or expired",
		})
	}

	resp, err := services.GetRefreshfeild(refreshtocken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err,
		})
	}

	if resp.RevokedAt != nil && time.Now().After(*resp.RevokedAt) {
		return c.Status(401).JSON(fiber.Map{
			"error": "Refresh token expired",
		})
	}

	NewAccesstoken, _ := services.GenarateJwt(resp.UserID)

	accessCookie := fiber.Cookie{
		Name:     "access_token",
		Value:    NewAccesstoken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&accessCookie)

	NewRefresh, _ := services.GenerateRefreshToken()
	//setting the reference cookie
	RefreshCookie := fiber.Cookie{
		Name:     "refresh_token",
		Value:    NewRefresh.Token,
		Expires:  NewRefresh.ExpiresAt,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&RefreshCookie) //setted here
	hashedtoocken := services.GenerateHash(NewRefresh.Token)
	services.SaveOrUpdateRefreshToken(db.DB, resp.UserID, hashedtoocken, NewRefresh.ExpiresAt)

	return c.Status(200).JSON(fiber.Map{
		"Refresh id ": refreshtocken,
		"found":       resp,
	})
}
