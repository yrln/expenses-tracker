package main

import (
	"log"

	"github.com/doctermath/expenses-tracker/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// load config
	cfg := config.Load()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Fatal(app.Listen(cfg.Addr))
}
