package handlers

import (
	"log"
	model "todo/models"
	"todo/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *fiber.Ctx) error {
	var user model.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Registration failed no matching data ",
		})
	}
	//setting random user_id
	user.User_id = uuid.NewString()
	//hashing pass
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("hashing failed ")
	}
	user.Password = string(hash)

	//sending the user for the insertion service
	res := services.InserUser(user)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to register user",
			"er":    res.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"User has been registered": user,
	})
}
