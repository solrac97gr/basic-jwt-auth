package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/basic-jwt-auth/config"
	"github.com/solrac97gr/basic-jwt-auth/handlers"
	"github.com/solrac97gr/basic-jwt-auth/middlewares"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)

	// Create a Login route
	app.Post("/login", handlers.Login)

	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)

	// Listen on port 3000
	app.Listen(":3000")
}
