package expenses

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo *Repository
}

func RegisterRoutes(r fiber.Router, db *sql.DB) {
	h := &Handler{
		repo: NewRepository(db),
	}

	r.Post("/expenses", h.create)
	r.Get("/expenses", h.list)
}

type createExpenseRequest struct {
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Note     string  `json:"note"`
}

func (h *Handler) create(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}
	userID, ok := userIDVal.(uint64)
	if !ok {
		// in case type assertion fails, try via string/int
		switch v := userIDVal.(type) {
		case int64:
			userID = uint64(v)
		case int:
			userID = uint64(v)
		case string:
			if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
				userID = parsed
			}
		}
		if userID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user id in context",
			})
		}
	}

	var req createExpenseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	req.Category = strings.TrimSpace(req.Category)
	req.Note = strings.TrimSpace(req.Note)

	if req.Amount <= 0 || req.Category == "" || req.Note == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "amount, category, and note are required",
		})
	}

	e, err := h.repo.Create(userID, req.Amount, req.Category, req.Note)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create expense",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(e)
}

func (h *Handler) list(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}
	userID, ok := userIDVal.(uint64)
	if !ok {
		switch v := userIDVal.(type) {
		case int64:
			userID = uint64(v)
		case int:
			userID = uint64(v)
		case string:
			if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
				userID = parsed
			}
		}
		if userID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user id in context",
			})
		}
	}

	categoryParam := strings.TrimSpace(c.Query("category", ""))
	var category *string
	if categoryParam != "" {
		category = &categoryParam
	}

	expenses, err := h.repo.ListByUser(userID, category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to list expenses",
		})
	}

	return c.JSON(expenses)
}
