package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yrln/expense-tracker/internal/auth"
	"github.com/yrln/expense-tracker/internal/config"
	"github.com/yrln/expense-tracker/internal/database"
)

func main() {
	// load config
	cfg := config.Load()

	db, err := database.OpenMySQL(cfg.DBDSN)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// public auth routes
	authGroup := app.Group("/auth")
	auth.RegisterRoutes(authGroup, db)

	log.Fatal(app.Listen(cfg.Addr))
}
