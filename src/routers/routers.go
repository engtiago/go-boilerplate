package routers

import (
	"github.com/engtiago/go-boilerplate/src/controllers"
	"github.com/engtiago/go-boilerplate/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	//login
	app.Post("/login", controllers.PostLogin)

	//users
	app.Post("/users", controllers.PostUser)

	app.Get("/users/:id", middleware.RequireAuth, controllers.GetUser)
	app.Get("/users", middleware.RequireAuth, controllers.GetUsers)
	app.Put("/users/:id", middleware.RequireAuth, controllers.PutUser)
	app.Delete("/users/:id", middleware.RequireAuth, controllers.DeleteUsers)
}
