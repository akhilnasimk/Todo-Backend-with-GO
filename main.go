package main

import (
	"fmt"
	"log"
	"todo/db"
	"todo/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	//db conection
	con := "host=localhost user=postgres password=123 dbname=todo port=5432 sslmode=disable"
	err := db.Init(con)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	fmt.Println(db.DB, "is ready")

	// fiber router setting
	route := fiber.New()

	route.Use(logger.New())
	route.Use(recover.New())

	route.Post("/Register", handlers.RegisterHandler)
	route.Post("/login", handlers.LoginHandler)
	// auth := route.Group("/auth")
	// auth.Use(Authmiddle())
	// {
	// 	auth.Get("/todo",GetAlltodo)
	// 	auth.Post("/todo",AddTodo)
	// 	auth.Patch("/todo/:id",PathTodo)
	// 	auth.Delete("/todo/:id",DeleteTodo)
	// }

	log.Fatal(route.Listen(":8080"))
}
