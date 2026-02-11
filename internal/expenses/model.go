package expenses

import "time"

type Expense struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
