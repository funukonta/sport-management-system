package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/utils"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}
