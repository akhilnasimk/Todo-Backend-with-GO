package handlers

import (
	"todo/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteTodo(c *fiber.Ctx) error {
	todo_id := c.Params("id")
	resp := services.Delete_todo(todo_id)
	if resp.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"Error ":   "faild to delete the todo ",
			"response": resp.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"Message": "todo has been deleted",
	})
}
