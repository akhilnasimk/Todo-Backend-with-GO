package main

import (
	"fmt"
	"log"
	"todo/db"
	"todo/handlers"
	"todo/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	route.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", 
		AllowCredentials: true,                    
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))
	route.Use(logger.New())
	route.Use(recover.New())

	route.Post("/Register", handlers.RegisterHandler)
	route.Post("/login", handlers.LoginHandler)
	auth := route.Group("/auth")
	auth.Use(middlewares.Authmiddle())
	{
		auth.Get("/todo", handlers.GetAlltodo)
		auth.Post("/todo", handlers.AddTodo)
		// auth.Patch("/todo/:id",PathTodo)
		auth.Delete("/todo/:id", handlers.DeleteTodo)
	}

	log.Fatal(route.Listen(":8080"))
}
