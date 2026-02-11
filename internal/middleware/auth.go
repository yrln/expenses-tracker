package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yrln/expense-tracker/internal/auth"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}

		tokenStr := parts[1]

		token, err := jwt.ParseWithClaims(tokenStr, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtSecret(), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		claims, ok := token.Claims.(*auth.Claims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}
