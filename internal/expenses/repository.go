package expenses

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(userID uint64, amount float64, category, note string) (Expense, error) {
	res, err := r.db.Exec(`
        INSERT INTO expenses (user_id, amount, category, note)
        VALUES (?, ?, ?, ?)
    `, userID, amount, category, note)
	if err != nil {
		return Expense{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Expense{}, err
	}

	var e Expense
	err = r.db.QueryRow(`
        SELECT id, user_id, amount, category, note, created_at
        FROM expenses
        WHERE id = ?
    `, id).Scan(&e.ID, &e.UserID, &e.Amount, &e.Category, &e.Note, &e.CreatedAt)
	if err != nil {
		return Expense{}, err
	}

	return e, nil
}

func (r *Repository) ListByUser(userID uint64, category *string) ([]Expense, error) {
	var rows *sql.Rows
	var err error

	if category != nil {
		rows, err = r.db.Query(`
            SELECT id, user_id, amount, category, note, created_at
            FROM expenses
            WHERE user_id = ? AND category = ?
            ORDER BY created_at DESC
        `, userID, *category)
	} else {
		rows, err = r.db.Query(`
            SELECT id, user_id, amount, category, note, created_at
            FROM expenses
            WHERE user_id = ?
            ORDER BY created_at DESC
        `, userID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Expense
	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.ID, &e.UserID, &e.Amount, &e.Category, &e.Note, &e.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, e)
	}

	return out, nil
}
