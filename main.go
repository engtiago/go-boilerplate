package main

import (
	"os"

	"github.com/engtiago/go-boilerplate/src/initializers"
	"github.com/engtiago/go-boilerplate/src/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routers.ApiRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(port)
}
