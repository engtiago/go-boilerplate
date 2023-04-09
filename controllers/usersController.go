package controllers

import (
	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type mdUser struct {
	Name     string
	Email    string
	Password string
}

func GetUser(c *fiber.Ctx) error {
	// TODO: Mudar isso depois para melhorar a esperiencia do usuário.
	//userrr := c.Locals("user")

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

	user.Password = ""

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

	for idx := range users {
		users[idx].Password = ""
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

	// Hash de password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return fiber.ErrBadGateway
	}

	// Create a user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return result.Error
	}
	user.Password = ""
	// Return it
	return c.JSON(user)
}

func PutUser(c *fiber.Ctx) error {
	// Get id
	id := c.Params("id")

	// Find user to update
	var user models.User
	resultFind := initializers.DB.Find(&user, id)
	if resultFind.Error != nil {
		return fiber.ErrBadRequest
	}

	// Get data form body
	body := mdUser{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	// Update a user
	resultUpdate := initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, Password: body.Password})

	if resultUpdate.Error != nil {
		return fiber.ErrBadRequest
	}
	user.Password = ""
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
