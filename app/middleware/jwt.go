package middleware

import (
	"github.com/directoryxx/auth-go/app/service"
	"github.com/gofiber/fiber/v2"

	jwt "github.com/golang-jwt/jwt/v4"
)

// JWTProtected - Protected protect routes
func JWTProtected(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		check := service.CheckLogin(claims["uuid"])
		if !check {
			c.Status(401)
			return c.JSON(fiber.Map{"error": "Unauthorized access"})
		}

		return c.Next()
	}
}
