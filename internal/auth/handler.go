package auth

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func RegisterRoutes(r fiber.Router, db *sql.DB) {
	h := &Handler{db: db}

	r.Post("/register", h.register)
	r.Post("/login", h.login)
}

func (h *Handler) register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "mortal kombat"})
}

func (h *Handler) login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "mortal kombat 1"})
}
