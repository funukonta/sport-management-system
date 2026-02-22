package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/utils"
)

type TeamRepository struct {
	db *sqlx.DB
}

func NewTeamRepository(db *sqlx.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) Create(team *models.Team) error {
	query := `INSERT INTO teams (name, logo, founded_year, address, city)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, team.Name, team.Logo, team.FoundedYear, team.Address, team.City).
		Scan(&team.ID, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrTeamExists)
		}

		return err
	}
	return nil
}

func (r *TeamRepository) FindAll() ([]models.Team, error) {
	var teams []models.Team
	query := `SELECT id, name, logo, founded_year, address, city, created_at, updated_at
			  FROM teams WHERE deleted_at IS NULL ORDER BY id`
	err := r.db.Select(&teams, query)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *TeamRepository) FindByID(id int) (*models.Team, error) {
	var team models.Team
	query := `SELECT id, name, logo, founded_year, address, city, created_at, updated_at
			  FROM teams WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&team, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewBadRequestError(ErrTeamNotFound)
		}
		return nil, err
	}
	return &team, nil
}

func (r *TeamRepository) Update(team *models.Team) error {
	query := `UPDATE teams SET name = $1, logo = $2, founded_year = $3, address = $4, city = $5, updated_at = NOW()
			  WHERE id = $6 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, team.Name, team.Logo, team.FoundedYear, team.Address, team.City, team.ID)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrTeamExists)
		}
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return utils.NewBadRequestError(ErrTeamNotFound)
	}
	return nil
}

func (r *TeamRepository) Delete(id int) error {
	query := `UPDATE teams SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return utils.NewBadRequestError(ErrTeamNotFound)
	}
	return nil
}
