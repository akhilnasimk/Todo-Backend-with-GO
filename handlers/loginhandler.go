package handlers

import (
	"time"
	"todo/db"
	model "todo/models"
	"todo/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *fiber.Ctx) error {
	var req model.Req

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": err,
		})
	}

	res, u := services.FindUser(req)

	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user not found",
			"val":   res.Error,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "wrong usre name or password ",
		})
	}

	//getting jwt tocken from jwt tofcken generator in service
	tokenString, err := services.GenarateJwt(u.User_id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error in tockrn checking": err,
		})
	}

	//getting refresh tocken
	referStruct, err := services.GenerateRefreshToken()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Failed to gnerate refresh tocken ": err,
		})
	}

	//setting access tocken
	accessCookie := fiber.Cookie{
		Name:     "access_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&accessCookie)

	//setting the reference cookie
	RefreshCookie := fiber.Cookie{
		Name:     "refresh_token",
		Value:    referStruct.Token,
		Expires:  referStruct.ExpiresAt,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	}
	c.Cookie(&RefreshCookie) //setted here
	//hashing password
	hashedtoocken := services.GenerateHash(referStruct.Token)
	errr := services.SaveOrUpdateRefreshToken(db.DB, u.User_id, hashedtoocken, referStruct.ExpiresAt) //running the db logic

	if errr != nil { //if something goes wrong 
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"error":    "failed to insert the refresh tocken into Db ",
			"response": errr,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"access tocken":   tokenString,
		"refresh tocken ": referStruct.Token,
	})
}
