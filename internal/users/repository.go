package users

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(email, passwordHash string) (User, error) {
	res, err := r.db.Exec(`
		INSERT INTO users (email, password_hash)
		VALUES (?, ?)
	`, email, passwordHash)

	if err != nil {
		return User{}, nil
	}

	id, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}

	var u User
	err = r.db.QueryRow(`
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE id = ?
	`, id).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)

	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *Repository) GetByEmail(email string) (User, error) {
	var u User
	err := r.db.QueryRow(`
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = ?
	`, email).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)

	if err != nil {
		return User{}, err
	}

	return u, nil
}
