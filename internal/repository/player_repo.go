package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/models"
	"github.com/sport-management-system/internal/utils"
)

type PlayerRepository struct {
	db *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) Create(player *models.Player) error {
	query := `INSERT INTO players (team_id, name, height, weight, position, jersey_number)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, player.TeamID, player.Name, player.Height, player.Weight, player.Position, player.JerseyNumber).Scan(&player.ID, &player.CreatedAt, &player.UpdatedAt)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrPlayerExist)
		}

		return err
	}
	return nil
}

func (r *PlayerRepository) FindAll() ([]models.Player, error) {
	var players []models.Player
	query := `SELECT id, team_id, name, height, weight, position, jersey_number, created_at, updated_at
			  FROM players WHERE deleted_at IS NULL ORDER BY id`
	err := r.db.Select(&players, query)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (r *PlayerRepository) FindAllPaginated(limit int, offset int) ([]models.Player, int, error) {
	var players []models.Player
	var total int

	// Get total count
	countQuery := `SELECT COUNT(*) FROM players WHERE deleted_at IS NULL`
	err := r.db.Get(&total, countQuery)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated data
	query := `SELECT id, team_id, name, height, weight, position, jersey_number, created_at, updated_at
			  FROM players WHERE deleted_at IS NULL ORDER BY id LIMIT $1 OFFSET $2`
	err = r.db.Select(&players, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return players, total, nil
}

func (r *PlayerRepository) FindByID(id int) (*models.Player, error) {
	var player models.Player
	query := `SELECT id, team_id, name, height, weight, position, jersey_number, created_at, updated_at
			  FROM players WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&player, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewBadRequestError(ErrPlayerNotFound)
		}
		return nil, err
	}
	return &player, nil
}

func (r *PlayerRepository) FindByTeamID(teamID int) ([]models.Player, error) {
	var players []models.Player
	query := `SELECT id, team_id, name, height, weight, position, jersey_number, created_at, updated_at
			  FROM players WHERE team_id = $1 AND deleted_at IS NULL ORDER BY jersey_number`
	err := r.db.Select(&players, query, teamID)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (r *PlayerRepository) Update(player *models.Player) error {
	query := `UPDATE players SET team_id = $1, name = $2, height = $3, weight = $4, position = $5,
			  jersey_number = $6, updated_at = NOW()
			  WHERE id = $7 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, player.TeamID, player.Name, player.Height, player.Weight, player.Position, player.JerseyNumber, player.ID)
	if err != nil {
		if utils.IsUniqueDataError(err) {
			return utils.NewBadRequestError(ErrPlayerExist)
		}

		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return utils.NewBadRequestError(ErrPlayerNotFound)
	}
	return nil
}

func (r *PlayerRepository) Delete(id int) error {
	query := `UPDATE players SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return utils.NewBadRequestError(ErrPlayerNotFound)
	}
	return nil
}
