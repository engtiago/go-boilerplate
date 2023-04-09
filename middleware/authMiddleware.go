package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/engtiago/go-boilerplate/initializers"
	"github.com/engtiago/go-boilerplate/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *fiber.Ctx) error {
	// Get token from headers
	headers := c.GetReqHeaders()

	tokenString := headers["Authorization"]

	if tokenString == "" {
		return fiber.ErrUnauthorized
	}

	splitToken := strings.Split(tokenString, "Bearer ")

	if splitToken[0] != "" {
		return fiber.ErrUnauthorized
	}

	fmt.Println(splitToken[0])

	token, _ := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expired token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return fiber.ErrUnauthorized
		}
		// find user subtoken
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return fiber.ErrUnauthorized
		}

		user.Password = ""

		// att to rec
		c.Locals("user", user)

		// continue
		return c.Next()
	} else {
		return fiber.ErrUnauthorized
	}

}
