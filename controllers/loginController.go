package controllers

import (
	"os"
	"time"

	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	result := initializers.DB.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		c.Status(400)
		return result.Error
	}

	if user.ID == 0 {
		return fiber.ErrBadRequest
	}

	// Compare password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return fiber.ErrBadRequest
	}

	// generate Jwt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}

	// Return it
	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
