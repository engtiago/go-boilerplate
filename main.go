package main

import (
	"os"

	"github.com/engtiago/go-boilerplate/controllers"
	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/middleware"
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

	app.Get("/users/:id", middleware.RequireAuth, controllers.GetUser)
	app.Get("/users", middleware.RequireAuth, controllers.GetUsers)
	app.Put("/users/:id", middleware.RequireAuth, controllers.PutUser)
	app.Delete("/users/:id", middleware.RequireAuth, controllers.DeleteUsers)

	port := os.Getenv("PORT")
	app.Listen(port)
}
