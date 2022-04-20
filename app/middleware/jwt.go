package middleware

import (
	"fmt"

	"github.com/directoryxx/auth-go/app/service"
	"github.com/gofiber/fiber/v2"

	jwt "github.com/golang-jwt/jwt/v4"
)

// JWTProtected - Protected protect routes
func JWTProtected(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return c.Next()
	}
}
