package main

import (
	"fiberProject/auth"
	"fiberProject/database"
	"fiberProject/handler"
	"fiberProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Login route
	app.Post("/login", auth.Login)

	// Accessible Routes
	app.Get("access", handler.Accessible)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	router.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
