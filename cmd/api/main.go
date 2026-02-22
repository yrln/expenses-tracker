package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yrln/expense-tracker/internal/auth"
	"github.com/yrln/expense-tracker/internal/config"
	"github.com/yrln/expense-tracker/internal/database"
	"github.com/yrln/expense-tracker/internal/expenses"
	"github.com/yrln/expense-tracker/internal/middleware"
)

func main() {
	// load config
	cfg := config.Load()

	db, err := database.OpenMySQL(cfg.DBDSN, cfg.DBCERT)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    "expenses-tracker",
			"env":     cfg.Env,
			"status":  "running",
			"version": "1.0.0",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service":    "expenses-tracker",
			"maintainer": "Yordan Flitz",
			"contact":    "doctermath@gmail.com",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	authGroup := app.Group("/auth")
	auth.RegisterRoutes(authGroup, db)

	api := app.Group("/api/v1", middleware.JWTAuth())
	auth.RegisterUserRoutes(api, db)
	expenses.RegisterRoutes(api, db)

	log.Fatal(app.Listen(cfg.Addr))
}
