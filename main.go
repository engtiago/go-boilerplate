package main

import (
	"os"

	"github.com/engtiago/go-boilerplate/controllers"
	"github.com/engtiago/go-boilerplate/initializers"
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

	//login
	app.Post("/login", controllers.PostLogin)

	//users
	app.Post("/users", controllers.PostUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Get("/users", controllers.GetUsers)
	app.Put("/users/:id", controllers.PutUser)
	app.Delete("/users/:id", controllers.DeleteUsers)

	port := os.Getenv("PORT")
	app.Listen(port)
}
