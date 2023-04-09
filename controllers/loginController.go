package controllers

import (
	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

type mdLogin struct {
	Email    string
	Password string
}

func PostLogin(c *fiber.Ctx) error {

	// Get data form body
	body := mdLogin{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	// Look for user
	var user models.User
	initializers.DB.First(&user, "email = ?", "")

	// Create a user
	//user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return result.Error
	}

	// Return it
	return c.JSON(user)
}
