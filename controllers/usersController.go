package controllers

import (
	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

type mdUser struct {
	Name     string
	Email    string
	Password string
}

func GetUser(c *fiber.Ctx) error {
	// Get id
	id := c.Params("id")

	// Get user
	var user models.User
	result := initializers.DB.Find(&user, id)

	if user.ID == 0 {
		return fiber.ErrNotFound
	}

	if result.Error != nil {
		c.Status(500)
		return result.Error
	}

	// Return it
	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	// Get users
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		return fiber.ErrBadGateway
	}

	// Return it
	return c.JSON(users)
}

func PostUser(c *fiber.Ctx) error {

	// Get data form body
	body := mdUser{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	// Create a user
	user := models.User{Name: body.Name, Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return result.Error
	}

	// Return it
	return c.JSON(user)
}

func UserPut(c *fiber.Ctx) error {
	// Get id
	id := c.Params("id")

	// Find user to update
	var user models.User
	resultFind := initializers.DB.Find(&user, id)
	if resultFind.Error != nil {
		return fiber.ErrNotFound
	}

	// Get data form body
	body := mdUser{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	// Update a user
	resultUpdate := initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, Password: body.Password})

	if resultUpdate.Error != nil {
		return fiber.ErrNotFound
	}

	// Return it
	return c.JSON(user)
}

func DeleteUsers(c *fiber.Ctx) error {
	// Get id
	id := c.Params("id")

	// find user
	var user models.User
	resultFind := initializers.DB.Find(&user, id)

	if resultFind.Error != nil {
		c.Status(500)
		return resultFind.Error
	}

	if user.ID == 0 {
		return fiber.ErrNotFound
	}

	// Get user
	resultDelete := initializers.DB.Delete(&models.User{}, id)

	if resultDelete.Error != nil {
		c.Status(500)
		return resultDelete.Error
	}

	// Return it
	c.Status(fiber.StatusNoContent)
	return nil
}
