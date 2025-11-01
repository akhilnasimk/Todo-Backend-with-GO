package handlers

import (
	"fmt"
	model "todo/models"

	"todo/services"

	"github.com/gofiber/fiber/v2"
)

func AddTodo(c *fiber.Ctx) error {
	var TodoReq model.Todo
	err := c.BodyParser(&TodoReq)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "bad request to the add todo ",
			"error":   err,
		})
	}
	user_id := c.Locals("user_id")
	TodoReq.UserID = user_id.(string)
	fmt.Println(TodoReq)
	res := services.Insert_todo(TodoReq)

	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message ": "failed to create todo ",
			"error":    res.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Request": TodoReq,
	})
}
